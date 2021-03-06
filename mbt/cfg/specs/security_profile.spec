ProtoObject: nwsec_pb2
Service: NwSecurity
enabled : True
dolEnabled : True
objects: 
    - object : 
        name : SecurityProfile
        key_handle : SecurityProfileKeyHandle
        max_objects : 255
        ignore:
            - op : GetAll
        create:
            api      : SecurityProfileCreate
            request  : SecurityProfileRequestMsg
            response : SecurityProfileResponseMsg
            pre_cb   : None
            post_cb  : callback://security_profile/PostCreateCb
        update:
            api      : SecurityProfileUpdate
            request  : SecurityProfileRequestMsg
            response : SecurityProfileResponseMsg
            pre_cb   : None
            post_cb  : callback://security_profile/PostUpdateCb
        delete:
            api      : SecurityProfileDelete
            request  : SecurityProfileDeleteRequestMsg
            response : SecurityProfileDeleteResponseMsg
            pre_cb   : None
            post_cb  : None
        get:
            api      : SecurityProfileGet
            request  : SecurityProfileGetRequestMsg
            response : SecurityProfileGetResponseMsg
            pre_cb   : None
            post_cb  : callback://security_profile/PostGetCb
    - object : 
        name : SecurityGroup
        key_handle : SecurityGroupKeyHandle
        ignore:
            - op : Get
            - op : Update
            - op : Delete
            - op : Create
            - op : GetAll
        ignore_v2:
            - op : GetAll
        create:
            api      : SecurityGroupCreate
            request  : SecurityGroupRequestMsg
            response : SecurityGroupRequestMsg
            pre_cb   : None
            post_cb  : None
        update:
            api      : SecurityGroupUpdate
            request  : SecurityGroupRequestMsg
            response : SecurityGroupRequestMsg
            pre_cb   : None
            post_cb  : None
        delete:
            api      : SecurityGroupDelete
            request  : SecurityGroupDeleteRequestMsg
            response : SecurityGroupDeleteResponseMsg
            pre_cb   : None
            post_cb  : None
        get:
            api      : SecurityGroupGet
            request  : SecurityGroupGetRequestMsg
            response : SecurityGroupGetResponseMsg
            pre_cb   : None
            post_cb  : None
    - object : 
        name : SecurityGroupPolicy
        key_handle : SecurityGroupPolicyKeyHandle
        ignore:
            - op : Get
            - op : Update
            - op : Delete
            - op : Create
            - op : GetAll
        create:
            api      : SecurityGroupPolicyCreate
            request  : SecurityGroupPolicyRequestMsg
            response : SecurityGroupPolicyRequestMsg
            pre_cb   : callback://security_profile/SecGroupPolicyPreCreateCb
            post_cb  : None
        update:
            api      : SecurityGroupPolicyUpdate
            request  : SecurityGroupPolicyRequestMsg
            response : SecurityGroupPolicyRequestMsg
            pre_cb   : callback://security_profile/SecGroupPolicyPreUpdateCb
            post_cb  : None
        delete:
            api      : SecurityGroupPolicyDelete
            request  : SecurityGroupPolicyDeleteRequestMsg
            response : SecurityGroupPolicyDeleteResponseMsg
            pre_cb   : None
            post_cb  : None
        get:
            api      : SecurityGroupPolicyGet
            request  : SecurityGroupPolicyGetRequestMsg
            response : SecurityGroupPolicyGetResponseMsg
            pre_cb   : None
            post_cb  : None
# DosPolicy is temporarily disabled because it doesn't have a KeyHandle
#    - object : 
#        name : DoSPolicy
#        ignore:
#        create:
#            api      : DoSPolicyCreate
#            request  : DoSPolicyRequestMsg
#            response : DoSPolicyRequestMsg
#            pre_cb   : None
#            post_cb  : None
#        update:
#            api      : DoSPolicyUpdate
#            request  : DoSPolicyRequestMsg
#            response : DoSPolicyRequestMsg
#            pre_cb   : None
#            post_cb  : None
#        delete:
#            api      : DoSPolicyDelete
#            request  : DoSPolicyDeleteRequestMsg
#            response : DoSPolicyDeleteResponseMsg
#            pre_cb   : None
#            post_cb  : None
#        get:
#            api      : DoSPolicyGet
#            request  : DoSPolicyGetRequestMsg
#            response : DoSPolicyGetResponseMsg
#            pre_cb   : None
#            post_cb  : None

