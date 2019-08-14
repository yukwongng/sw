#ifndef __ENG_IF_HPP__
#define __ENG_IF_HPP__

#include <map>
#include <openssl/ssl.h>
#include <openssl/engine.h>
#include <openssl/evp.h>
#include "logger.hpp"
#include "nic/include/base.hpp"
#include "pse_intf.h"
#include "crypto_asym.hpp"

using namespace std;

namespace eng_if {

extern ENGINE   *eng_if_engine;
extern BIO      *eng_if_bio;

typedef EVP_MD eng_evp_md_t;

/*
 * Engine generic key
 */
class dole_key_t
{
public:
    dole_key_t() :
        key_type_(EVP_PKEY_NONE),
        key_idx_(CRYPTO_ASYM_KEY_IDX_INVALID),
        n_(nullptr),
        d_e_(nullptr),
        digest_padded_(nullptr),
        wait_for_completion_(false)
    {
    }

    dole_key_t&
    key_type(int key_type)
    {
        key_type_ = key_type;
        return *this;
    }
    dole_key_t&
    key_idx(crypto_asym::key_idx_t key_idx)
    {
        key_idx_ = key_idx;
        return *this;
    }
    dole_key_t&
    n(dp_mem_t *n)
    {
        n_ = n;
        return *this;
    }
    dole_key_t&
    d_e(dp_mem_t *d_e)
    {
        d_e_ = d_e;
        return *this;
    }
    dole_key_t&
    digest_padded(dp_mem_t *digest_padded)
    {
        digest_padded_ = digest_padded;
        return *this;
    }
    dole_key_t&
    wait_for_completion(bool wait_for_completion)
    {
        wait_for_completion_ = wait_for_completion;
        return *this;
    }

    int key_type(void) { return key_type_; }
    crypto_asym::key_idx_t key_idx(void) { return key_idx_; }
    dp_mem_t *n(void) { return n_; }
    dp_mem_t *d_e(void) { return d_e_; }
    dp_mem_t *digest_padded(void) { return digest_padded_; }
    bool wait_for_completion(void) { return wait_for_completion_; }

private:
    int                         key_type_;
    crypto_asym::key_idx_t      key_idx_;
    dp_mem_t                    *n_;            // RSA modulus
    dp_mem_t                    *d_e_;          // RSA public or private exponent
    dp_mem_t                    *digest_padded_;
    bool                        wait_for_completion_;
};

/*
 * Engine digest parameters
 */
class digest_params_t
{
public:
    digest_params_t() :
        msg_(nullptr),
        digest_(nullptr)
    {
    }

    digest_params_t&
    hash_algo(const string& hash_algo)
    {
        hash_algo_ = hash_algo;
        return *this;
    }
    digest_params_t&
    msg(dp_mem_t *msg)
    {
        msg_ = msg;
        return *this;
    }
    digest_params_t&
    digest(dp_mem_t *digest)
    {
        digest_ = digest;
        return *this;
    }

    const string& hash_algo(void) { return hash_algo_; }
    dp_mem_t *msg(void) { return msg_; }
    dp_mem_t *digest(void) { return digest_; }

private:
    string                      hash_algo_;
    dp_mem_t                    *msg_;
    dp_mem_t                    *digest_;
};

/*
 * Engine RSA sign parameters
 */
class rsa_sign_params_t
{
public:
    rsa_sign_params_t() :
        md_(nullptr),
        pad_mode_(RSA_NO_PADDING),
        key_idx_(CRYPTO_ASYM_KEY_IDX_INVALID),
        n_(nullptr),
        d_e_(nullptr),
        digest_(nullptr),
        digest_padded_(nullptr),
        sig_actual_(nullptr),
        rsa_(nullptr),
        failure_expected_(false),
        wait_for_completion_(false)
    {
    }

    rsa_sign_params_t&
    md(const eng_evp_md_t *md)
    {
        md_ = md;
        return *this;
    }
    rsa_sign_params_t&
    pad_mode(int pad_mode)
    {
        pad_mode_ = pad_mode;
        return *this;
    }
    rsa_sign_params_t&
    key_idx(crypto_asym::key_idx_t key_idx)
    {
        key_idx_ = key_idx;
        return *this;
    }
    rsa_sign_params_t&
    n(dp_mem_t *n)
    {
        n_ = n;
        return *this;
    }
    rsa_sign_params_t&
    d_e(dp_mem_t *d_e)
    {
        d_e_ = d_e;
        return *this;
    }
    rsa_sign_params_t&
    digest(dp_mem_t *digest)
    {
        digest_ = digest;
        return *this;
    }
    rsa_sign_params_t&
    digest_padded(dp_mem_t *digest_padded)
    {
        digest_padded_ = digest_padded;
        return *this;
    }
    rsa_sign_params_t&
    sig_actual(dp_mem_t *sig_actual)
    {
        sig_actual_ = sig_actual;
        return *this;
    }
    rsa_sign_params_t&
    rsa(void *rsa)
    {
        rsa_ = rsa;
        return *this;
    }
    rsa_sign_params_t&
    failure_expected(bool failure_expected)
    {
        failure_expected_ = failure_expected;
        return *this;
    }
    rsa_sign_params_t&
    wait_for_completion(bool wait_for_completion)
    {
        wait_for_completion_ = wait_for_completion;
        return *this;
    }

    const eng_evp_md_t *md(void) { return md_; }
    int pad_mode(void) { return pad_mode_; }
    crypto_asym::key_idx_t key_idx(void) { return key_idx_; }
    dp_mem_t *n(void) { return n_; }
    dp_mem_t *d_e(void) { return d_e_; }
    dp_mem_t *digest(void) { return digest_; }
    dp_mem_t *digest_padded(void) { return digest_padded_; }
    dp_mem_t *sig_actual(void) { return sig_actual_; }
    void *rsa(void) { return rsa_; }
    bool failure_expected(void) { return failure_expected_; }
    bool wait_for_completion(void) { return wait_for_completion_; }

private:
    const eng_evp_md_t          *md_;
    int                         pad_mode_;
    crypto_asym::key_idx_t      key_idx_;
    dp_mem_t                    *n_;            // RSA modulus
    dp_mem_t                    *d_e_;          // RSA public or private exponent
    dp_mem_t                    *digest_;
    dp_mem_t                    *digest_padded_;
    dp_mem_t                    *sig_actual_;
    void                        *rsa_;
    bool                        failure_expected_;
    bool                        wait_for_completion_;
};

/*
 * Engine RSA verify parameters
 */
class rsa_verify_params_t
{
public:
    rsa_verify_params_t() :
        md_(nullptr),
        pad_mode_(RSA_NO_PADDING),
        key_idx_(CRYPTO_ASYM_KEY_IDX_INVALID),
        n_(nullptr),
        d_e_(nullptr),
        sig_expected_(nullptr),
        digest_(nullptr),
        digest_padded_(nullptr),
        rsa_(nullptr),
        failure_expected_(false),
        wait_for_completion_(false)
    {
    }

    rsa_verify_params_t&
    md(const eng_evp_md_t *md)
    {
        md_ = md;
        return *this;
    }
    rsa_verify_params_t&
    pad_mode(int pad_mode)
    {
        pad_mode_ = pad_mode;
        return *this;
    }
    rsa_verify_params_t&
    key_idx(crypto_asym::key_idx_t key_idx)
    {
        key_idx_ = key_idx;
        return *this;
    }
    rsa_verify_params_t&
    n(dp_mem_t *n)
    {
        n_ = n;
        return *this;
    }
    rsa_verify_params_t&
    d_e(dp_mem_t *d_e)
    {
        d_e_ = d_e;
        return *this;
    }
    rsa_verify_params_t&
    sig_expected(dp_mem_t *sig_expected)
    {
        sig_expected_ = sig_expected;
        return *this;
    }
    rsa_verify_params_t&
    digest(dp_mem_t *digest)
    {
        digest_ = digest;
        return *this;
    }
    rsa_verify_params_t&
    digest_padded(dp_mem_t *digest_padded)
    {
        digest_padded_ = digest_padded;
        return *this;
    }
    rsa_verify_params_t&
    rsa(void *rsa)
    {
        rsa_ = rsa;
        return *this;
    }
    rsa_verify_params_t&
    failure_expected(bool failure_expected)
    {
        failure_expected_ = failure_expected;
        return *this;
    }
    rsa_verify_params_t&
    wait_for_completion(bool wait_for_completion)
    {
        wait_for_completion_ = wait_for_completion;
        return *this;
    }

    const eng_evp_md_t *md(void) { return md_; }
    int pad_mode(void) { return pad_mode_; }
    crypto_asym::key_idx_t key_idx(void) { return key_idx_; }
    dp_mem_t *n(void) { return n_; }
    dp_mem_t *d_e(void) { return d_e_; }
    dp_mem_t *sig_expected(void) { return sig_expected_; }
    dp_mem_t *digest(void) { return digest_; }
    dp_mem_t *digest_padded(void) { return digest_padded_; }
    void *rsa(void) { return rsa_; }
    bool failure_expected(void) { return failure_expected_; }
    bool wait_for_completion(void) { return wait_for_completion_; }

private:
    const eng_evp_md_t          *md_;
    int                         pad_mode_;
    crypto_asym::key_idx_t      key_idx_;
    dp_mem_t                    *n_;            // RSA modulus
    dp_mem_t                    *d_e_;          // RSA public or private exponent
    dp_mem_t                    *sig_expected_;
    dp_mem_t                    *digest_;
    dp_mem_t                    *digest_padded_;
    void                        *rsa_;
    bool                        failure_expected_;
    bool                        wait_for_completion_;
};

/*
 * Engine EC domain parameters
 */
class ec_domain_params_t
{
public:
    ec_domain_params_t() :
        curve_nid_(0),
        d_(nullptr),
        p_(nullptr),
        n_(nullptr),
        xg_(nullptr),
        yg_(nullptr),
        a_(nullptr),
        b_(nullptr),
        da_(nullptr)
    {
    }

    ec_domain_params_t&
    curve_name(string& curve_name)
    {
        curve_name_ = curve_name;
        return *this;
    }
    ec_domain_params_t&
    curve_nid(int curve_nid)
    {
        curve_nid_ = curve_nid;
        return *this;
    }
    ec_domain_params_t&
    d(dp_mem_t *d)
    {
        d_ = d;
        return *this;
    }
    ec_domain_params_t&
    p(dp_mem_t *p)
    {
        p_ = p;
        return *this;
    }
    ec_domain_params_t&
    n(dp_mem_t *n)
    {
        n_ = n;
        return *this;
    }
    ec_domain_params_t&
    xg(dp_mem_t *xg)
    {
        xg_ = xg;
        return *this;
    }
    ec_domain_params_t&
    yg(dp_mem_t *yg)
    {
        yg_ = yg;
        return *this;
    }
    ec_domain_params_t&
    a(dp_mem_t *a)
    {
        a_ = a;
        return *this;
    }
    ec_domain_params_t&
    b(dp_mem_t *b)
    {
        b_ = b;
        return *this;
    }
    ec_domain_params_t&
    da(dp_mem_t *da)
    {
        da_ = da;
        return *this;
    }

    string& curve_name(void) { return curve_name_; }
    int curve_nid(void) { return curve_nid_; }
    dp_mem_t *d(void) { return d_; }
    dp_mem_t *p(void) { return p_; }
    dp_mem_t *n(void) { return n_; }
    dp_mem_t *xg(void) { return xg_; }
    dp_mem_t *yg(void) { return yg_; }
    dp_mem_t *a(void) { return a_; }
    dp_mem_t *b(void) { return b_; }
    dp_mem_t *da(void) { return da_; }

private:
    string                      curve_name_;
    int                         curve_nid_;
    dp_mem_t                    *d_;
    dp_mem_t                    *p_;
    dp_mem_t                    *n_;
    dp_mem_t                    *xg_;
    dp_mem_t                    *yg_;
    dp_mem_t                    *a_;
    dp_mem_t                    *b_;
    dp_mem_t                    *da_;
};

/*
 * Engine EC signature generation parameters
 */
class ec_sign_params_t
{
public:
    ec_sign_params_t() :
        curve_nid_(0),
        md_(nullptr),
        ec_(nullptr),
        key_idx_(CRYPTO_ASYM_KEY_IDX_INVALID),
        k_random_(nullptr),
        digest_(nullptr),
        sig_output_vec_(nullptr),
        skip_DER_encode_(false),
        failure_expected_(false),
        wait_for_completion_(false)
    {
    }

    ec_sign_params_t&
    curve_nid(int curve_nid)
    {
        curve_nid_ = curve_nid;
        return *this;
    }
    ec_sign_params_t&
    md(const eng_evp_md_t *md)
    {
        md_ = md;
        return *this;
    }
    ec_sign_params_t&
    ec(void *ec)
    {
        ec_ = ec;
        return *this;
    }
    ec_sign_params_t&
    key_idx(crypto_asym::key_idx_t key_idx)
    {
        key_idx_ = key_idx;
        return *this;
    }
    ec_sign_params_t&
    k_random(dp_mem_t *k_random)
    {
        k_random_ = k_random;
        return *this;
    }
    ec_sign_params_t&
    digest(dp_mem_t *digest)
    {
        digest_ = digest;
        return *this;
    }
    ec_sign_params_t&
    sig_output_vec(dp_mem_t *sig_output_vec)
    {
        sig_output_vec_ = sig_output_vec;
        return *this;
    }
    ec_sign_params_t&
    skip_DER_encode(bool skip_DER_encode)
    {
        skip_DER_encode_ = skip_DER_encode;
        return *this;
    }
    ec_sign_params_t&
    failure_expected(bool failure_expected)
    {
        failure_expected_ = failure_expected;
        return *this;
    }
    ec_sign_params_t&
    wait_for_completion(bool wait_for_completion)
    {
        wait_for_completion_ = wait_for_completion;
        return *this;
    }

    int curve_nid(void) { return curve_nid_; }
    const eng_evp_md_t *md(void) { return md_; }
    void *ec(void) { return ec_; }
    crypto_asym::key_idx_t key_idx(void) { return key_idx_; }
    dp_mem_t *k_random(void) { return k_random_; }
    dp_mem_t *digest(void) { return digest_; }
    dp_mem_t *sig_output_vec(void) { return sig_output_vec_; }
    bool skip_DER_encode(void) { return skip_DER_encode_; }
    bool failure_expected(void) { return failure_expected_; }
    bool wait_for_completion(void) { return wait_for_completion_; }

private:
    int                         curve_nid_;
    const eng_evp_md_t          *md_;
    void                        *ec_;
    crypto_asym::key_idx_t      key_idx_;
    dp_mem_t                    *k_random_;
    dp_mem_t                    *digest_;
    dp_mem_t                    *sig_output_vec_;
    bool                        skip_DER_encode_;  // skip Distinguished Encoding Rules
    bool                        failure_expected_;
    bool                        wait_for_completion_;
};

/*
 * Engine EC signature verification parameters
 */
class ec_verify_params_t
{
public:
    ec_verify_params_t() :
        curve_nid_(0),
        md_(nullptr),
        ec_(nullptr),
        key_idx_(CRYPTO_ASYM_KEY_IDX_INVALID),
        digest_(nullptr),
        sig_input_vec_(nullptr),
        sig_r_(nullptr),
        sig_s_(nullptr),
        skip_DER_decode_(false),
        failure_expected_(false),
        wait_for_completion_(false)
    {
    }

    ec_verify_params_t&
    curve_nid(int curve_nid)
    {
        curve_nid_ = curve_nid;
        return *this;
    }
    ec_verify_params_t&
    md(const eng_evp_md_t *md)
    {
        md_ = md;
        return *this;
    }
    ec_verify_params_t&
    ec(void *ec)
    {
        ec_ = ec;
        return *this;
    }
    ec_verify_params_t&
    key_idx(crypto_asym::key_idx_t key_idx)
    {
        key_idx_ = key_idx;
        return *this;
    }
    ec_verify_params_t&
    digest(dp_mem_t *digest)
    {
        digest_ = digest;
        return *this;
    }
    ec_verify_params_t&
    sig_input_vec(dp_mem_t *sig_input_vec)
    {
        sig_input_vec_ = sig_input_vec;
        return *this;
    }
    ec_verify_params_t&
    sig_r(dp_mem_t *sig_r)
    {
        sig_r_ = sig_r;
        return *this;
    }
    ec_verify_params_t&
    sig_s(dp_mem_t *sig_s)
    {
        sig_s_ = sig_s;
        return *this;
    }
    ec_verify_params_t
    skip_DER_decode(bool skip_DER_decode)
    {
        skip_DER_decode_ = skip_DER_decode;
        return *this;
    }
    ec_verify_params_t&
    failure_expected(bool failure_expected)
    {
        failure_expected_ = failure_expected;
        return *this;
    }
    ec_verify_params_t&
    wait_for_completion(bool wait_for_completion)
    {
        wait_for_completion_ = wait_for_completion;
        return *this;
    }

    int curve_nid(void) { return curve_nid_; }
    const eng_evp_md_t *md(void) { return md_; }
    void *ec(void) { return ec_; }
    crypto_asym::key_idx_t key_idx(void) { return key_idx_; }
    dp_mem_t *digest(void) { return digest_; }
    dp_mem_t *sig_input_vec(void) { return sig_input_vec_; }
    dp_mem_t *sig_r(void) { return sig_r_; }
    dp_mem_t *sig_s(void) { return sig_s_; }
    bool skip_DER_decode(void) { return skip_DER_decode_; }
    bool failure_expected(void) { return failure_expected_; }
    bool wait_for_completion(void) { return wait_for_completion_; }

private:
    int                         curve_nid_;
    const eng_evp_md_t          *md_;
    void                        *ec_;
    crypto_asym::key_idx_t      key_idx_;
    dp_mem_t                    *digest_;
    dp_mem_t                    *sig_input_vec_;
    dp_mem_t                    *sig_r_;        // (r, s) are fragments of sig_input_vec
    dp_mem_t                    *sig_s_;
    bool                        skip_DER_decode_;  // skip Distinguished Encoding Rules
    bool                        failure_expected_;
    bool                        wait_for_completion_;
};

bool init(const char *engine_path);

bool dp_mem_pad_in_place(dp_mem_t *msg,
                         uint32_t to_len);
bool dp_mem_unpad_in_place(dp_mem_t *msg,
                           uint32_t to_len);
bool dp_mem_trunc_in_place(dp_mem_t *msg,
                           uint32_t bits_len);
bool bn_to_dp_mem_pad(const BIGNUM *bn,
                      dp_mem_t *mem,
                      int to_len);
bool bn_to_dp_mem(const BIGNUM *bn,
                  dp_mem_t *mem);
BIGNUM *dp_mem_to_bn(dp_mem_t *mem);

const eng_evp_md_t *hash_algo_find(const string& hash_algo);
const eng_evp_md_t *digest_gen(digest_params_t& params);

bool rsa_sign(rsa_sign_params_t& params);
bool rsa_verify(rsa_verify_params_t& params);

bool ec_domain_params_gen(ec_domain_params_t& params);
bool ec_sign(ec_sign_params_t& params);
bool ec_verify(ec_verify_params_t& params);

} // namespace eng_if

#endif /* __ENG_IF_HPP__ */

