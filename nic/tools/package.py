import os
from subprocess import call
import pdb
import sys

output_dir = 'fake_root_target'
arm_pkg    = 1

files = [ 'nic/tools/pack_local_files.txt' ]

if len(sys.argv) == 2:
    if sys.argv[1] == 'x86_64':
        arm_pkg = 0
        files.append('nic/tools/pack_local_files_x86_64.txt')
        print ("packaging x86_64")
    else:
        print ("packaging aarch64")
        files.append('nic/tools/pack_local_files_aarch64.txt')
else:
    print ("packaging aarch64")
    files.append('nic/tools/pack_local_files_aarch64.txt')

for input_file in files:
    print ("Processing input file: " + input_file)
    f  = open(input_file, 'r')

    for line in f:
        items = line.split()
        directory = output_dir + '/' + items[1]
        if not os.path.exists(directory):
            print ('Creating dir: ' + directory)
            os.makedirs(directory)

        if (os.path.isdir(items[0])):
            cmd = 'cp -rL ' + items[0] + '/* ' + directory
            call(cmd, shell=True)
        else:
            cmd = 'cp ' + items[0] + ' ' + directory
            print (cmd)
            call(cmd, shell=True)

# strip the libs and binaries only for arm packaging
if arm_pkg == 1:
    for root, dirs, files in os.walk(output_dir):
        for file in files:
            if '.so' in file or 'hal' in file or 'linkmgr' in file:
                non_stripped = os.path.join(root, file)
                call(['chmod', '755', non_stripped])
                call(['/tool/toolchain/aarch64-1.1/bin/aarch64-linux-gnu-objcopy', '--only-keep-debug', non_stripped, non_stripped + '.debug'])
                call(['/tool/toolchain/aarch64-1.1/bin/aarch64-linux-gnu-strip', non_stripped])
                call(['/tool/toolchain/aarch64-1.1/bin/aarch64-linux-gnu-objcopy', '--add-gnu-debuglink=' + non_stripped + '.debug', non_stripped])

cmd = 'mkdir -p ' + output_dir + '/nic/lib'
call(cmd, shell=True)

# remove dol plugin for aarch64
if arm_pkg == 1:
    cmd = 'rm -rf ' + output_dir + '/nic/conf/plugins/dol'
    call(cmd, shell=True)

# remove *.log from nic/conf/init_bins libs
cmd = 'find ' + output_dir + '/nic/conf/init_bins -name "*.log" | xargs rm -f'
call(cmd, shell=True)

##### TODO REVERT LATER #####

# remove *.a from platform libs
cmd = 'rm -f ' + output_dir + '/platform/lib/*.a'
call(cmd, shell=True)

# remove csrlite until main csr lib is not removed
cmd = 'rm -f ' + output_dir + '/nic/lib/libcsrlite.so'
call(cmd, shell=True)

# rename libzmq.so to libzmq.so.3
if arm_pkg == 0:
    cmd = 'mv ' + output_dir + '/nic/lib/libzmq.so ' + output_dir + '/nic/lib/libzmq.so.3'
    call(cmd, shell=True)

cmd = 'cd ' + output_dir + ' && tar --exclude=*.debug -cf ../nic/nic.tar *'
call(cmd, shell=True)

# create tar.gz
cmd = 'cd ' + output_dir + ' && tar --exclude=*.debug -czf ../nic/nic.tgz * && chmod 766 ../nic/nic.tgz'
call(cmd, shell=True)
