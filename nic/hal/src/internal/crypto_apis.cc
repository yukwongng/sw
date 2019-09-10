//-----------------------------------------------------------------------------
// {C} Copyright 2017 Pensando Systems Inc. All rights reserved
//-----------------------------------------------------------------------------

#include "nic/include/base.hpp"
#include "nic/hal/hal.hpp"
#include "nic/hal/iris/include/hal_state.hpp"
#include "nic/include/pd_api.hpp"
#include "nic/sdk/include/sdk/crypto_apis.hpp"
#include "nic/hal/src/internal/crypto_cert_store.hpp"
// #include "nic/hal/pd/capri/capri_barco_asym_apis.hpp"
// #include "nic/hal/pd/capri/capri_barco_sym_apis.hpp"
#include <openssl/pem.h>

namespace hal {


/* Crypto Cert Store related APIs */

void *
crypto_cert_store_get_key_func(void *entry)
{
    SDK_ASSERT(entry != NULL);
    return (void *)&(((crypto_cert_t *)entry)->cert_id);
}

uint32_t crypto_cert_store_key_size(void) {
    return sizeof(crypto_cert_id_t);
}

/* Asym APIs */
hal_ret_t crypto_asym_api_ecc_point_mul(internal::CryptoApiRequest &req,
        internal::CryptoApiResponse *resp)
{
    hal_ret_t           ret = HAL_RET_OK;
    uint32_t            key_size;
    uint8_t             px[ECC_MAX_KEY_SIZE];
    uint8_t             py[ECC_MAX_KEY_SIZE];
    pd::pd_capri_barco_asym_ecc_point_mul_args_t args;
    pd::pd_func_args_t          pd_func_args = {0};

    key_size = req.ecc_point_mul_fp().ecc_domain_params().keysize();

    switch (key_size) {
        case 28: /* P-224 */
        case 32: /* P-256 */
        case 48: /* P-384 */
        case 66: /* P-521 */
            args.key_size = key_size;
            args.p = (uint8_t *)req.ecc_point_mul_fp().ecc_domain_params().p().data();
            args.n = (uint8_t *)req.ecc_point_mul_fp().ecc_domain_params().n().data();
            args.xg = (uint8_t *)req.ecc_point_mul_fp().ecc_domain_params().g().x().data();
            args.yg = (uint8_t *)req.ecc_point_mul_fp().ecc_domain_params().g().y().data();
            args.a = (uint8_t *)req.ecc_point_mul_fp().ecc_domain_params().a().data();
            args.b = (uint8_t *)req.ecc_point_mul_fp().ecc_domain_params().b().data();
            args.x1 = (uint8_t *)req.ecc_point_mul_fp().ecc_point().x().data();
            args.y1 = (uint8_t *)req.ecc_point_mul_fp().ecc_point().y().data();
            args.k = (uint8_t *)req.ecc_point_mul_fp().k().data();
            args.x3 = px;
            args.y3 = py;
            pd_func_args.pd_capri_barco_asym_ecc_point_mul = &args;
            ret = pd::hal_pd_call(pd::PD_FUNC_ID_BARCO_ASYM_ECC_MUL, &pd_func_args);
            if (ret == HAL_RET_OK) {
                resp->mutable_ecc_point_mul_fp()->mutable_q()->mutable_x()->assign(
                        (const char*)px, (size_t) key_size);
                resp->mutable_ecc_point_mul_fp()->mutable_q()->mutable_y()->assign(
                        (const char*)py, (size_t) key_size);
            }
            break;
        default:
            HAL_TRACE_ERR("Unsupported key size: {}", key_size);
            ret = HAL_RET_INVALID_ARG;
            break;
    }
    if (ret == HAL_RET_OK) {
        resp->set_api_status(types::API_STATUS_OK);
    }
    else {
        resp->set_api_status(types::API_STATUS_ERR);
    }

    return ret;
}


hal_ret_t crypto_asym_api_ecdsa_sig_gen(internal::CryptoApiRequest &req,
        internal::CryptoApiResponse *resp)
{
    hal_ret_t           ret = HAL_RET_OK;
    uint32_t            key_size;
    uint8_t             r[ECC_MAX_KEY_SIZE];
    uint8_t             s[ECC_MAX_KEY_SIZE];
    pd::pd_capri_barco_asym_ecdsa_p256_sig_gen_args_t args;
    pd::pd_func_args_t          pd_func_args = {0};

    key_size = req.ecdsa_sig_gen_fp().ecc_domain_params().keysize();

    switch (key_size) {
        case 32:
            args.key_idx = req.ecdsa_sig_gen_fp().key_idx();
            args.p = (uint8_t *)req.ecdsa_sig_gen_fp().ecc_domain_params().p().data();
            args.n = (uint8_t *)req.ecdsa_sig_gen_fp().ecc_domain_params().n().data();
            args.xg = (uint8_t *)req.ecdsa_sig_gen_fp().ecc_domain_params().g().x().data();
            args.yg = (uint8_t *)req.ecdsa_sig_gen_fp().ecc_domain_params().g().y().data();
            args.a = (uint8_t *)req.ecdsa_sig_gen_fp().ecc_domain_params().a().data();
            args.b = (uint8_t *)req.ecdsa_sig_gen_fp().ecc_domain_params().b().data();
            args.da = (uint8_t *)req.ecdsa_sig_gen_fp().da().data();
            args.k = (uint8_t *)req.ecdsa_sig_gen_fp().k().data();
            args.h = (uint8_t *)req.ecdsa_sig_gen_fp().h().data();
            args.r = r;
            args.s = s;
            args.async_en = req.ecdsa_sig_gen_fp().async_en();

            pd_func_args.pd_capri_barco_asym_ecdsa_p256_sig_gen = &args;
            ret = pd::hal_pd_call(pd::PD_FUNC_ID_BARCO_ASYM_ECDSA_P256_SIG_GEN, &pd_func_args);
            if (ret == HAL_RET_OK) {
                resp->mutable_ecdsa_sig_gen_fp()->mutable_r()->assign(
                        (const char*)r, (size_t) key_size);
                resp->mutable_ecdsa_sig_gen_fp()->mutable_s()->assign(
                        (const char*)s, (size_t) key_size);
            }
            break;
        default:
            HAL_TRACE_ERR("Unsupported key size: {}", key_size);
            ret = HAL_RET_INVALID_ARG;
            break;
    }
    if (ret == HAL_RET_OK) {
        resp->set_api_status(types::API_STATUS_OK);
    }
    else {
        resp->set_api_status(types::API_STATUS_ERR);
    }
    return ret;
}

hal_ret_t crypto_asym_api_ecdsa_sig_verify(internal::CryptoApiRequest &req,
        internal::CryptoApiResponse *resp)
{
    hal_ret_t           ret = HAL_RET_OK;
    uint32_t            key_size;
    pd::pd_capri_barco_asym_ecdsa_p256_sig_verify_args_t args;
    pd::pd_func_args_t          pd_func_args = {0};

    key_size = req.ecdsa_sig_verify_fp().ecc_domain_params().keysize();

    switch (key_size) {
        case 32:
            args.p = (uint8_t *)req.ecdsa_sig_verify_fp().ecc_domain_params().p().data();
            args.n = (uint8_t *)req.ecdsa_sig_verify_fp().ecc_domain_params().n().data();
            args.xg = (uint8_t *)req.ecdsa_sig_verify_fp().ecc_domain_params().g().x().data();
            args.yg = (uint8_t *)req.ecdsa_sig_verify_fp().ecc_domain_params().g().y().data();
            args.a = (uint8_t *)req.ecdsa_sig_verify_fp().ecc_domain_params().a().data();
            args.b = (uint8_t *)req.ecdsa_sig_verify_fp().ecc_domain_params().b().data();
            args.xq = (uint8_t *)req.ecdsa_sig_verify_fp().q().x().data();
            args.yq = (uint8_t *)req.ecdsa_sig_verify_fp().q().y().data();
            args.r = (uint8_t *)req.ecdsa_sig_verify_fp().r().data();
            args.s = (uint8_t *)req.ecdsa_sig_verify_fp().s().data();
            args.h = (uint8_t *)req.ecdsa_sig_verify_fp().h().data();
            args.async_en = (uint8_t *)req.ecdsa_sig_verify_fp().async_en();
            pd_func_args.pd_capri_barco_asym_ecdsa_p256_sig_verify = &args;
            ret = pd::hal_pd_call(pd::PD_FUNC_ID_BARCO_ASYM_ECDSA_P256_SIG_VER, &pd_func_args);
            break;
        default:
            HAL_TRACE_ERR("Unsupported key size: {}", key_size);
            ret = HAL_RET_INVALID_ARG;
            break;
    }
    if (ret == HAL_RET_OK) {
        resp->set_api_status(types::API_STATUS_OK);
    }
    else {
        resp->set_api_status(types::API_STATUS_ERR);
    }

    return ret;
}

hal_ret_t crypto_asym_api_rsa_encrypt(internal::CryptoApiRequest &req,
        internal::CryptoApiResponse *resp)
{
    hal_ret_t           ret = HAL_RET_OK;
    uint32_t            key_size;
    uint8_t             cipher_text[RSA_MAX_KEY_SIZE];
    pd::pd_capri_barco_asym_rsa2k_encrypt_args_t args;
    pd::pd_func_args_t          pd_func_args = {0};

    key_size = req.rsa_encrypt().keysize();

    switch (key_size) {
        case 256:
            args.n = (uint8_t *)req.rsa_encrypt().mod_n().data();
            args.e = (uint8_t *)req.rsa_encrypt().e().data();
            args.m = (uint8_t *)req.rsa_encrypt().plain_text().data();
            args.c = cipher_text;
            args.async_en = req.rsa_encrypt().async_en();
            pd_func_args.pd_capri_barco_asym_rsa2k_encrypt = &args;
            ret = pd::hal_pd_call(pd::PD_FUNC_ID_BARCO_ASYM_RSA2K_ENCRYPT, &pd_func_args);
            break;
        default:
            HAL_TRACE_ERR("Unsupported key size: {}", key_size);
            ret = HAL_RET_INVALID_ARG;
            break;
    }
    if (ret == HAL_RET_OK) {
        resp->mutable_rsa_encrypt()->mutable_cipher_text()->assign(
                (const char*)cipher_text, (size_t) key_size);
        resp->set_api_status(types::API_STATUS_OK);
    }
    else {
        resp->set_api_status(types::API_STATUS_ERR);
    }

    return ret;
}

hal_ret_t crypto_asym_api_rsa_decrypt(internal::CryptoApiRequest &req,
        internal::CryptoApiResponse *resp)
{
    hal_ret_t           ret = HAL_RET_OK;
    uint32_t            key_size;
    uint8_t             plain_text[RSA_MAX_KEY_SIZE];
    pd::pd_capri_barco_asym_rsa2k_decrypt_args_t args;
    pd::pd_func_args_t  pd_func_args = {0};

    key_size = req.rsa_decrypt().keysize();

    switch (key_size) {
        case 256:
            args.n = (uint8_t *)req.rsa_decrypt().mod_n().data();
            args.d =  (uint8_t *)req.rsa_decrypt().d().data();
            args.c = (uint8_t *)req.rsa_decrypt().cipher_text().data();
            args.m = plain_text;
            pd_func_args.pd_capri_barco_asym_rsa2k_decrypt = &args;
            ret = pd::hal_pd_call(pd::PD_FUNC_ID_BARCO_ASYM_RSA2K_DECRYPT, &pd_func_args);
            break;
        default:
            HAL_TRACE_ERR("Unsupported key size: {}", key_size);
            ret = HAL_RET_INVALID_ARG;
            break;
    }
    if (ret == HAL_RET_OK) {
        resp->mutable_rsa_decrypt()->mutable_plain_text()->assign(
                (const char*)plain_text, (size_t) key_size);
        resp->set_api_status(types::API_STATUS_OK);
    }
    else {
        resp->set_api_status(types::API_STATUS_ERR);
    }

    return ret;
}

hal_ret_t crypto_asym_api_rsa_crt_decrypt(internal::CryptoApiRequest &req,
        internal::CryptoApiResponse *resp)
{
    hal_ret_t           ret = HAL_RET_OK;
    uint32_t            key_size;
    uint8_t             plain_text[RSA_MAX_KEY_SIZE];
    pd::pd_capri_barco_asym_rsa2k_crt_decrypt_args_t args;
    pd::pd_func_args_t          pd_func_args = {0};

    key_size = req.rsa_crt_decrypt().keysize();

    switch (key_size) {
        case 256:
            args.key_idx = req.rsa_crt_decrypt().key_idx();
            args.p = (uint8_t *)req.rsa_crt_decrypt().p().data();
            args.q = (uint8_t *)req.rsa_crt_decrypt().q().data();
            args.dp = (uint8_t *)req.rsa_crt_decrypt().dp().data();
            args.dq = (uint8_t *)req.rsa_crt_decrypt().dq().data();
            args.qinv = (uint8_t *)req.rsa_crt_decrypt().qinv().data();
            args.c = (uint8_t *)req.rsa_crt_decrypt().cipher_text().data();
            args.m = plain_text;
            args.async_en = req.rsa_crt_decrypt().async_en();
            pd_func_args.pd_capri_barco_asym_rsa2k_crt_decrypt = &args;
            ret = pd::hal_pd_call(pd::PD_FUNC_ID_BARCO_ASYM_RSA2K_CRT_DECRYPT, &pd_func_args);
            break;
        default:
            HAL_TRACE_ERR("Unsupported key size: {}", key_size);
            ret = HAL_RET_INVALID_ARG;
            break;
    }
    if (ret == HAL_RET_OK) {
        resp->mutable_rsa_crt_decrypt()->mutable_plain_text()->assign(
                (const char*)plain_text, (size_t) key_size);
        resp->set_api_status(types::API_STATUS_OK);
    }
    else {
        resp->set_api_status(types::API_STATUS_ERR);
    }

    return ret;
}

static hal_ret_t
crypto_asym_api_setup_ec_priv_key(EVP_PKEY *pkey,
                                  internal::CryptoAsymApiRespSetupPrivateKey *setup_key_resp)
{
    hal_ret_t           ret = HAL_RET_OK;
    EC_KEY              *ec_key = NULL;
    const EC_GROUP      *group = NULL;
    const BIGNUM        *priv_key = NULL;
    const EC_POINT      *ec_point = NULL;
    BIGNUM              *p = NULL, *order = NULL;
    BIGNUM              *xg = NULL, *yg = NULL;
    BIGNUM              *a = NULL, *b = NULL;
    BN_CTX              *ctx = NULL;
    uint8_t             buf_p[256] = {0}, buf_n[256] = {0};
    uint8_t             buf_xg[256] = {0}, buf_yg[256] = {0};
    uint8_t             buf_a[256] = {0}, buf_b[256] = {0};
    uint8_t             buf_da[256] = {0};
    int32_t             key_idx = -1;
    pd::pd_capri_barco_asym_ecdsa_p256_setup_priv_key_args_t args = {0};
    pd::pd_func_args_t          pd_func_args = {0};

    if(!pkey)
        return HAL_RET_INVALID_ARG;

    ec_key = EVP_PKEY_get0_EC_KEY(pkey);
    if(!ec_key) {
        HAL_TRACE_ERR("Failed to get ec key");
        return HAL_RET_ERR;
    }

    group = EC_KEY_get0_group(ec_key);
    priv_key = EC_KEY_get0_private_key(ec_key);
    if(group == NULL || priv_key == NULL) {
        HAL_TRACE_ERR("Failed to get the group/private key from the key");
        return HAL_RET_ERR;
    }

    if((ec_point = EC_GROUP_get0_generator(group)) == NULL) {
        HAL_TRACE_ERR("Failed to retrive ec_point");
        return HAL_RET_ERR;
    }

    if((ctx = BN_CTX_new()) == NULL) {
        HAL_TRACE_ERR("Failed to allocate BN ctx");
        goto cleanup;
    }

    BN_CTX_start(ctx);
    p = BN_CTX_get(ctx);
    a = BN_CTX_get(ctx);
    b = BN_CTX_get(ctx);
    xg = BN_CTX_get(ctx);
    yg = BN_CTX_get(ctx);
    order = BN_CTX_get(ctx);

    if(order == NULL) {
        HAL_TRACE_ERR("Failed to allocate memory for p, a, b etc");
        ret = HAL_RET_OOM;
        goto cleanup;
    }

    if(!EC_GROUP_get_order(group, order, ctx)) {
        HAL_TRACE_ERR("Failed to get order from the group");
        ret = HAL_RET_ERR;
        goto cleanup;
    }

    if (EC_METHOD_get_field_type(EC_GROUP_method_of(group))
                            == NID_X9_62_prime_field) {
        HAL_TRACE_DEBUG("key field type primefield");
        if ((!EC_GROUP_get_curve_GFp(group, p, a, b, ctx)) ||
            (!EC_POINT_get_affine_coordinates_GFp(group, ec_point,
                                                  xg, yg, ctx))) {
            HAL_TRACE_ERR("Failed to get curve params for prime field");
            ret = HAL_RET_ERR;
            goto cleanup;
        }
    } else {
        if ((!EC_GROUP_get_curve_GF2m(group, p, a, b, ctx)) ||
            (!EC_POINT_get_affine_coordinates_GF2m(group, ec_point,
                                                   xg, yg, ctx))) {
            HAL_TRACE_ERR("Failed to get curve params for binary field");
            ret = HAL_RET_ERR;
            goto cleanup;
        }
    }

    BN_bn2bin(p, buf_p);
    BN_bn2bin(order, buf_n);
    BN_bn2bin(xg, buf_xg);
    BN_bn2bin(yg, buf_yg);
    BN_bn2bin(a, buf_a);
    BN_bn2bin(b, buf_b);
    BN_bn2bin(priv_key, buf_da);

    args.p = buf_p;
    args.n = buf_n;
    args.xg = buf_xg;
    args.yg = buf_yg;
    args.a = buf_a;
    args.b = buf_b;
    args.da = buf_da;
    args.key_idx = &key_idx;

    pd_func_args.pd_capri_barco_asym_ecdsa_p256_setup_priv_key = &args;
    ret = pd::hal_pd_call(pd::PD_FUNC_ID_BARCO_ASYM_ECDSA_P256_SETUP_PRIV_KEY,
                          &pd_func_args);
    HAL_TRACE_DEBUG("Received key_idx: {}", key_idx);
    setup_key_resp->set_key_type(types::CRYPTO_ASYM_KEY_TYPE_ECDSA);
    setup_key_resp->mutable_ecdsa_key_info()->set_sign_key_idx(key_idx);

cleanup:
    if(ctx) {
        BN_CTX_end(ctx);
        BN_CTX_free(ctx);
    }

    return ret;
}

static hal_ret_t
crypto_asym_api_setup_rsa_sig_gen_priv_key(RSA *rsa,
                                           int32_t &key_idx)
{
    BIGNUM              *n = NULL, *e = NULL, *d = NULL;
    uint8_t             buf_n[256] = {0}, buf_d[256] = {0};
    pd::pd_capri_barco_asym_rsa2k_setup_sig_gen_priv_key_args_t args = {0};
    pd::pd_func_args_t          pd_func_args = {0};

    if(!rsa) {
        return HAL_RET_INVALID_ARG;
    }

    // Extract params
    RSA_get0_key(rsa,
                 (const BIGNUM**)&n,
                 (const BIGNUM**)&e,
                 (const BIGNUM**)&d);

    BN_bn2binpad(n, buf_n, 256);
    BN_bn2binpad(d, buf_d,256);

    args.n = (uint8_t *)buf_n;
    args.d = (uint8_t *)buf_d;
    args.key_idx = &key_idx;
    pd_func_args.pd_capri_barco_asym_rsa2k_setup_sig_gen_priv_key = &args;
    return pd::hal_pd_call(pd::PD_FUNC_ID_ASYM_RSA2K_SETUP_SIG_GEN_PRIV_KEY, &pd_func_args);
}

static hal_ret_t
crypto_asym_api_setup_rsa_decrypt_priv_key(RSA *rsa,
                                           int32_t &key_idx)
{
    BIGNUM              *p = NULL, *q = NULL, *dp = NULL, *dq = NULL, *qinv = NULL;
    uint8_t             buf_p[256] = {0}, buf_q[256] = {0};
    uint8_t             buf_dp[256] = {0}, buf_dq[256] = {0}, buf_qinv[256] = {0};
    pd::pd_capri_barco_asym_rsa2k_crt_setup_decrypt_priv_key_args_t args = {0};
    pd::pd_func_args_t          pd_func_args = {0};

    if(!rsa) {
        return HAL_RET_INVALID_ARG;
    }

    // Extract params
    RSA_get0_factors(rsa,
                     (const BIGNUM**)&p,
                     (const BIGNUM**)&q);
    RSA_get0_crt_params(rsa,
                        (const BIGNUM**)&dp,
                        (const BIGNUM**)&dq,
                        (const BIGNUM**)&qinv);

    BN_bn2binpad(p, buf_p, 256);
    BN_bn2binpad(q, buf_q, 256);
    BN_bn2binpad(dp, buf_dp, 256);
    BN_bn2binpad(dq, buf_dq, 256);
    BN_bn2binpad(qinv, buf_qinv, 256);

    args.p = (uint8_t *)buf_p;
    args.q = (uint8_t *)buf_q;
    args.dp = (uint8_t *)buf_dp;
    args.dq = (uint8_t *)buf_dq;
    args.qinv = (uint8_t *)buf_qinv;
    args.key_idx = &key_idx;
    pd_func_args.pd_capri_barco_asym_rsa2k_crt_setup_decrypt_priv_key = &args;
    return pd::hal_pd_call(pd::PD_FUNC_ID_ASYM_RSA2K_CRT_SETUP_DECRYPT_PRIV_KEY, &pd_func_args);
}

static hal_ret_t
crypto_asym_api_setup_rsa_priv_key(EVP_PKEY *pkey,
                                   internal::CryptoAsymApiRespSetupPrivateKey *setup_key_resp)
{
    hal_ret_t           ret = HAL_RET_OK;
    RSA                 *rsa = NULL;
    uint32_t            key_size = 0;
    int32_t             sig_gen_key_idx = -1;
    int32_t             decrypt_key_idx = -1;

    if(!pkey) {
        return HAL_RET_INVALID_ARG;
    }

    // get rsa
    rsa = EVP_PKEY_get0_RSA(pkey);
    if(!rsa) {
        HAL_TRACE_ERR("Failed to extract RSA from key");
        return HAL_RET_ERR;
    }

    key_size = RSA_size(rsa);
    if(key_size != 256) {
        HAL_TRACE_ERR("Invalid key size: {}", key_size);
        return HAL_RET_INVALID_ARG;
    }

    // setup sig gen key
    ret = crypto_asym_api_setup_rsa_sig_gen_priv_key(rsa, sig_gen_key_idx);
    if(ret != HAL_RET_OK) {
        HAL_TRACE_ERR("Failed to setup rsa sig gen key: {}", ret);
        goto cleanup;
    }

    ret = crypto_asym_api_setup_rsa_decrypt_priv_key(rsa, decrypt_key_idx);
    if(ret != HAL_RET_OK) {
        HAL_TRACE_ERR("Failed to setup rsa decrypt key: {}", ret);
        goto cleanup;
    }

    HAL_TRACE_DEBUG("Received sig_gen key: {}, decrypt key: {}",
                    sig_gen_key_idx, decrypt_key_idx);
    setup_key_resp->set_key_type(types::CRYPTO_ASYM_KEY_TYPE_RSA);
    setup_key_resp->mutable_rsa_key_info()->set_sign_key_idx(sig_gen_key_idx);
    setup_key_resp->mutable_rsa_key_info()->set_decrypt_key_idx(decrypt_key_idx);

    return ret;
cleanup:
    //TODO: remove allocated keys
    return ret;
}
hal_ret_t crypto_asym_api_setup_priv_key(internal::CryptoApiRequest &req,
					 internal::CryptoApiResponse *resp)
{
    hal_ret_t           ret = HAL_RET_OK;
    BIO                 *bio = NULL;
    EVP_PKEY            *pkey = NULL;
    uint32_t            key_type = 0;

    // decode the key
    bio = BIO_new_mem_buf(req.setup_priv_key().key().c_str(), -1);
    if(!bio) {
        HAL_TRACE_ERR("Failed to allocate bio");
        ret = HAL_RET_ERR;
        goto end;
    }

    // decode the key
    pkey = PEM_read_bio_PrivateKey(bio, NULL, 0, NULL);
    if(!pkey) {
        HAL_TRACE_ERR("Failed to decode the key");
        ret = HAL_RET_ERR;
        goto end;
    }

    // Program the key in hw
    key_type = EVP_PKEY_base_id(pkey);

    HAL_TRACE_DEBUG("Received privkey type {}", key_type);
    switch (key_type) {
    case EVP_PKEY_EC:
        ret = crypto_asym_api_setup_ec_priv_key(pkey, resp->mutable_setup_priv_key());
        break;
    case EVP_PKEY_RSA:
        ret = crypto_asym_api_setup_rsa_priv_key(pkey, resp->mutable_setup_priv_key());
        break;
    default:
        HAL_TRACE_ERR("Unsupported key type: {}", key_type);
        ret = HAL_RET_INVALID_ARG;
        break;
    }
end:
    if (ret == HAL_RET_OK) {
        resp->set_api_status(types::API_STATUS_OK);
    }
    else {
        resp->set_api_status(types::API_STATUS_ERR);
    }
    return ret;
}

static hal_ret_t
crypto_asym_api_setup_rsa_priv_key_ex(internal::CryptoApiRequest &req,
                                   internal::CryptoApiResponse *resp)
{
    hal_ret_t                   ret = HAL_RET_OK;
    uint32_t                    key_size;
    int32_t                     key_idx;
    pd::pd_capri_barco_asym_rsa_setup_priv_key_args_t args = {0};
    pd::pd_func_args_t          pd_func_args = {0};

    key_size = req.setup_priv_key_ex().rsa_key().key_size();

    if ((req.setup_priv_key_ex().rsa_key().n().length() != key_size) ||
            (req.setup_priv_key_ex().rsa_key().e().length() != key_size) ||
            (req.setup_priv_key_ex().rsa_key().d().length() != key_size)) {

        HAL_TRACE_ERR("Length mismatch in the input key");
        ret = HAL_RET_ERR;
        goto end;
    }

    args.key_size = key_size;
    args.n = (uint8_t *)req.setup_priv_key_ex().rsa_key().n().c_str();
    args.d = (uint8_t *)req.setup_priv_key_ex().rsa_key().d().c_str();
    args.key_idx = &key_idx;
    pd_func_args.pd_capri_barco_asym_rsa_setup_priv_key = &args;
    ret = pd::hal_pd_call(pd::PD_FUNC_ID_ASYM_RSA_SETUP_PRIV_KEY, &pd_func_args);

end:
    if (ret == HAL_RET_OK) {
        resp->mutable_setup_priv_key_ex()->set_key_type(types::CRYPTO_ASYM_KEY_TYPE_RSA);
        resp->mutable_setup_priv_key_ex()->set_key_idx(key_idx);
        resp->set_api_status(types::API_STATUS_OK);
    }
    else {
        resp->set_api_status(types::API_STATUS_ERR);
    }
    return ret;
}

hal_ret_t crypto_asym_api_setup_priv_key_ex(internal::CryptoApiRequest &req,
					 internal::CryptoApiResponse *resp)
{
    hal_ret_t           ret = HAL_RET_OK;
    uint32_t            key_type = 0;

    switch (req.setup_priv_key_ex().key_type()) {
        case types::CRYPTO_ASYM_KEY_TYPE_ECDSA:
#if 0
            ret = crypto_asym_api_setup_ec_priv_key(pkey, resp->mutable_setup_priv_key());
#endif
            break;
        case types::CRYPTO_ASYM_KEY_TYPE_RSA:
            ret = crypto_asym_api_setup_rsa_priv_key_ex(req, resp);
            break;
        default:
            HAL_TRACE_ERR("Unsupported key type: {}", key_type);
            ret = HAL_RET_INVALID_ARG;
            break;
    }
    return ret;
}

hal_ret_t crypto_asym_api_rsa_sig_gen(internal::CryptoApiRequest &req,
        internal::CryptoApiResponse *resp)
{
    hal_ret_t           ret = HAL_RET_OK;
    uint32_t            key_size;
    uint8_t             s[RSA_MAX_KEY_SIZE];
    pd::pd_capri_barco_asym_rsa_sig_gen_args_t args;
    pd::pd_func_args_t          pd_func_args = {0};

    key_size = req.rsa_sig_gen().keysize();

    switch (key_size) {
        case 128:   // 1K
        case 256:   // 2K
        case 384:   // 3K
        case 512:   // 4K
            args.key_size = key_size;
            args.key_idx = req.rsa_sig_gen().key_idx();
            args.n = (uint8_t *)req.rsa_sig_gen().mod_n().data();
            args.d = (uint8_t *)req.rsa_sig_gen().d().data();
            args.h = (uint8_t *)req.rsa_sig_gen().h().data();
            args.s = s;
            args.async_en = req.rsa_sig_gen().async_en();
            pd_func_args.pd_capri_barco_asym_rsa_sig_gen = &args;
            ret = pd::hal_pd_call(pd::PD_FUNC_ID_BARCO_ASYM_RSA_SIG_GEN, &pd_func_args);
            break;
        default:
            HAL_TRACE_ERR("Unsupported key size: {}", key_size);
            ret = HAL_RET_INVALID_ARG;
            break;
    }
    if (ret == HAL_RET_OK) {
        resp->mutable_rsa_sig_gen()->mutable_s()->assign(
                (const char*)s, (size_t) key_size);
        resp->set_api_status(types::API_STATUS_OK);
    }
    else {
        resp->set_api_status(types::API_STATUS_ERR);
    }

    return ret;
}

hal_ret_t crypto_asym_api_rsa_sig_verify(internal::CryptoApiRequest &req,
        internal::CryptoApiResponse *resp)
{
    hal_ret_t           ret = HAL_RET_OK;
    uint32_t            key_size;
    pd::pd_capri_barco_asym_rsa2k_sig_verify_args_t args;
    pd::pd_func_args_t          pd_func_args = {0};

    key_size = req.rsa_sig_verify().keysize();

    switch (key_size) {
        case 256:
            args.n = (uint8_t *)req.rsa_sig_verify().mod_n().data();
            args.e = (uint8_t *)req.rsa_sig_verify().e().data();
            args.h = (uint8_t *)req.rsa_sig_verify().h().data();
            args.s = (uint8_t *)req.rsa_sig_verify().s().data();
            pd_func_args.pd_capri_barco_asym_rsa2k_sig_verify = &args;
            ret = pd::hal_pd_call(pd::PD_FUNC_ID_BARCO_ASYM_RSA2K_SIG_VERIFY, &pd_func_args);
            break;
        default:
            HAL_TRACE_ERR("Unsupported key size: {}", key_size);
            ret = HAL_RET_INVALID_ARG;
            break;
    }
    if (ret == HAL_RET_OK) {
        resp->set_api_status(types::API_STATUS_OK);
    }
    else {
        resp->set_api_status(types::API_STATUS_ERR);
    }

    return ret;
}

static hal_ret_t
crypto_asym_extract_cert_pubkey_params(crypto_cert_t *cert)
{
    hal_ret_t     ret = HAL_RET_OK;
    EVP_PKEY      *pkey = NULL;
    RSA           *rsa = NULL;
    BIGNUM        *n = NULL, *e = NULL, *d = NULL;
    const EC_KEY  *eckey = NULL;

    if(!cert || !cert->x509_cert)
        return HAL_RET_INVALID_ARG;

    pkey = X509_get0_pubkey(cert->x509_cert);
    cert->pub_key.key_type = EVP_PKEY_base_id(pkey);
    HAL_TRACE_DEBUG("Received pubkey type {}", cert->pub_key.key_type);

    switch(cert->pub_key.key_type) {
    case EVP_PKEY_RSA:
        rsa = EVP_PKEY_get0_RSA(pkey);
        if(!rsa) {
            HAL_TRACE_ERR("Failed to extract rsa from the cert");
            ret = HAL_RET_ERR;
            goto end;
        }
        cert->pub_key.key_len = RSA_size(rsa);
        RSA_get0_key(rsa, (const BIGNUM**)&n, (const BIGNUM**)&e, (const BIGNUM **)&d);
        cert->pub_key.u.rsa_params.mod_n_len = BN_num_bytes(n);
        BN_bn2bin(n, cert->pub_key.u.rsa_params.mod_n);

        cert->pub_key.u.rsa_params.e_len = BN_num_bytes(e);
        BN_bn2bin(e, cert->pub_key.u.rsa_params.e);

        HAL_TRACE_DEBUG("n: {}", cert->pub_key.u.rsa_params.mod_n_len);
        HAL_TRACE_DEBUG("e: {}", cert->pub_key.u.rsa_params.e_len);
        break;

    case EVP_PKEY_EC:
        eckey = EVP_PKEY_get0_EC_KEY(pkey);
        if(!eckey) {
            HAL_TRACE_ERR("Failed to extract eckey from the cert");
            ret = HAL_RET_ERR;
            goto end;
        }
        cert->pub_key.u.ec_params.group = EC_KEY_get0_group(eckey);
        cert->pub_key.u.ec_params.point = EC_KEY_get0_public_key(eckey);
        break;

    default:
        HAL_TRACE_ERR("Invalid Key type {}", cert->pub_key.key_type);
        break;
    }
 end:

    return ret;
}
hal_ret_t crypto_asym_api_setup_cert(internal::CryptoApiRequest &req,
                                     internal::CryptoApiResponse *resp)
{
    hal_ret_t           ret = HAL_RET_OK;
    crypto_cert_t       *cert = NULL;
    sdk_ret_t           sdk_ret;
    BIO                 *bio = NULL;
    if(req.setup_cert().cert_id() <= 0) {
        HAL_TRACE_ERR("Invalid cert-id: {}",
                      req.setup_cert().cert_id());
        ret = HAL_RET_INVALID_ARG;
        goto end;
    }

    //find cert with the id
    cert = find_cert_by_id(req.setup_cert().cert_id());

    if(req.setup_cert().update_type() == internal::ADD_UPDATE) {
        if(!cert) {
            // Allocate a new cert
            cert = crypto_cert_alloc_init();
            if(!cert) {
                HAL_TRACE_ERR("Failed to allocate cert");
                ret = HAL_RET_OOM;
                goto end;
            }

            cert->cert_id = req.setup_cert().cert_id();

            // Add the cert to the hashtable
            sdk_ret = g_hal_state->crypto_cert_store_id_ht()->
                    insert(cert, &cert->ht_ctxt);
            if(sdk_ret != sdk::SDK_RET_OK) {
                HAL_TRACE_ERR("Failed to add cert to the hashtable");
                ret = hal_sdk_ret_to_hal_ret(sdk_ret);
                goto end;
            } // add hashtable
        } // Add new cert case
        // update body
        cert = find_cert_by_id(req.setup_cert().cert_id());
        if(!cert) {
            HAL_TRACE_ERR("cert cannot be found {}",
                          req.setup_cert().cert_id());
            SDK_ASSERT(0);
        }
        bio = BIO_new_mem_buf(req.setup_cert().body().c_str(), -1);
        if(!bio) {
            HAL_TRACE_ERR("Failed to allocate bio");
            ret = HAL_RET_ERR;
            goto end;
        }

        // Decode the cert
        cert->x509_cert = PEM_read_bio_X509(bio, NULL, 0, NULL);
        if(!cert->x509_cert) {
            HAL_TRACE_ERR("Failed to decode the certificate");
            ret = HAL_RET_ERR;
            goto end;
        }

        // Extract public key params
        ret = crypto_asym_extract_cert_pubkey_params(cert);
        if(ret != HAL_RET_OK) {
            HAL_TRACE_ERR("Failed to extract pubkey params");
            goto end;
        }

        cert->next_cert_id = req.setup_cert().next_cert_id();
        HAL_TRACE_DEBUG("Updated cert with idx: {}", cert->cert_id);
    } else {
        // Delete case
        if(!cert) {
            HAL_TRACE_DEBUG("Cert already deleted");
            ret = HAL_RET_OK;
            goto end;
        }
        // delete is from the hashtable
        g_hal_state->crypto_cert_store_id_ht()->remove(&cert->cert_id);

        // free memory
        crypto_cert_free(cert);
    }

end:
    if(ret == HAL_RET_OK) {
        resp->set_api_status(types::API_STATUS_OK);
    } else {
        resp->set_api_status(types::API_STATUS_ERR);
        if(cert) {
            crypto_cert_free(cert);
        }
    }
    if(bio) {
        BIO_free(bio);
    }
    return ret;
}


hal_ret_t crypto_symm_api_hash_request(internal::CryptoApiRequest &req,
				       internal::CryptoApiResponse *resp,
				       bool generate)
{
    hal_ret_t                     ret = HAL_RET_OK;
    int32_t                       digest_len, exp_digest_len;
    uint8_t                       digest[CRYPTO_MAX_HASH_DIGEST_LEN];
    internal::CryptoApiHashType hashtype;
    pd::pd_capri_barco_sym_hash_process_request_args_t args;
    pd::pd_func_args_t          pd_func_args = {0};

    if (generate) {
        hashtype = req.hash_generate().hashtype();
	digest_len = req.hash_generate().digest_len();
    } else {
        hashtype = req.hash_verify().hashtype();
	digest_len = req.hash_verify().digest_len();
    }

    switch(hashtype) {
    case internal::CRYPTOAPI_HASHTYPE_SHA1:
    case internal::CRYPTOAPI_HASHTYPE_HMAC_SHA1:
        exp_digest_len = CRYPTO_SHA1_DIGEST_LEN;
        break;
    case internal::CRYPTOAPI_HASHTYPE_SHA224:
    case internal::CRYPTOAPI_HASHTYPE_HMAC_SHA224:
        exp_digest_len = CRYPTO_SHA224_DIGEST_LEN;
        break;
    case internal::CRYPTOAPI_HASHTYPE_SHA256:
    case internal::CRYPTOAPI_HASHTYPE_HMAC_SHA256:
        exp_digest_len = CRYPTO_SHA256_DIGEST_LEN;
        break;
    case internal::CRYPTOAPI_HASHTYPE_SHA384:
    case internal::CRYPTOAPI_HASHTYPE_HMAC_SHA384:
        exp_digest_len = CRYPTO_SHA384_DIGEST_LEN;
        break;
    case internal::CRYPTOAPI_HASHTYPE_SHA512:
    case internal::CRYPTOAPI_HASHTYPE_HMAC_SHA512:
        exp_digest_len = CRYPTO_SHA512_DIGEST_LEN;
        break;
    case internal::CRYPTOAPI_HASHTYPE_SHA3_224:
        exp_digest_len = CRYPTO_SHA3_224_DIGEST_LEN;
        break;
    case internal::CRYPTOAPI_HASHTYPE_SHA3_256:
        exp_digest_len = CRYPTO_SHA3_256_DIGEST_LEN;
        break;
    case internal::CRYPTOAPI_HASHTYPE_SHA3_384:
        exp_digest_len = CRYPTO_SHA3_384_DIGEST_LEN;
        break;
    case internal::CRYPTOAPI_HASHTYPE_SHA3_512:
        exp_digest_len = CRYPTO_SHA3_512_DIGEST_LEN;
        break;
    default:
        HAL_TRACE_ERR("Unsupported Hash type: {}",
		      CryptoApiHashType_Name(hashtype));
        resp->set_api_status(types::API_STATUS_ERR);
	return HAL_RET_INVALID_ARG;
    }

    if (exp_digest_len > digest_len) {
        HAL_TRACE_ERR("Digest length invalid: {}",
		      CryptoApiHashType_Name(hashtype));
        resp->set_api_status(types::API_STATUS_ERR);
	return HAL_RET_INVALID_ARG;
    }

    args.hash_type = (CryptoApiHashType) hashtype;
    args.generate = generate;
    args.key = generate ? (uint8_t *)req.hash_generate().key().data() : (uint8_t *)req.hash_verify().key().data();
    args.key_len = generate ? (uint32_t)req.hash_generate().key_len() : (uint32_t)req.hash_verify().key_len();
    args.data = generate ? (uint8_t *)req.hash_generate().data().data() : (uint8_t *)req.hash_verify().data().data();
    args.data_len = generate ? (uint32_t)req.hash_generate().data_len() : (uint32_t)req.hash_verify().data_len();
    args.output_digest = generate ? digest : (uint8_t *)req.hash_verify().digest().data();
    args.digest_len = digest_len;
    pd_func_args.pd_capri_barco_sym_hash_process_request = &args;
    ret = pd::hal_pd_call(pd::PD_FUNC_ID_BARCO_SYM_HASH_PROC_REQ, &pd_func_args);
    if (ret == HAL_RET_OK) {
        if (generate) {
	    resp->mutable_hash_generate()->mutable_digest()->assign(
		    (const char*)digest, (size_t) exp_digest_len);
	}
        resp->set_api_status(types::API_STATUS_OK);
    }
    else {
        resp->set_api_status(types::API_STATUS_ERR);
    }

    return ret;
}

hal_ret_t crypto_asym_api_fips_rsa_sig_gen(internal::CryptoApiRequest &req,
        internal::CryptoApiResponse *resp)
{
    hal_ret_t           ret = HAL_RET_OK;
    uint32_t            key_size;
    uint8_t             s[RSA_MAX_KEY_SIZE];
    pd::pd_capri_barco_asym_fips_rsa_sig_gen_args_t args;
    pd::pd_func_args_t          pd_func_args = {0};

    key_size = req.fips_rsa_sig_gen().mod_n().size();

    switch (key_size) {
        case 128:   // 1K
        case 192:   // 1.5K
        case 256:   // 2K
        case 384:   // 3K
        case 512:   // 4K
            args.key_size = key_size;
            args.key_idx = req.fips_rsa_sig_gen().key_idx();
            args.n = (uint8_t *)req.fips_rsa_sig_gen().mod_n().data();
            args.e = (uint8_t *)req.fips_rsa_sig_gen().e().data();
            args.msg = (uint8_t *)req.fips_rsa_sig_gen().msg().data();
            args.msg_len = req.fips_rsa_sig_gen().msg().size();
            args.s = s;
            args.hash_type = req.fips_rsa_sig_gen().hash_type();
            args.sig_scheme = req.fips_rsa_sig_gen().sig_scheme();

            args.async_en = false;
            pd_func_args.pd_capri_barco_asym_fips_rsa_sig_gen = &args;
            ret = pd::hal_pd_call(pd::PD_FUNC_ID_BARCO_ASYM_FIPS_RSA_SIG_GEN, &pd_func_args);
            break;
        default:
            HAL_TRACE_ERR("Unsupported key size: {}", key_size);
            ret = HAL_RET_INVALID_ARG;
            break;
    }
    if (ret == HAL_RET_OK) {
        resp->mutable_fips_rsa_sig_gen()->mutable_s()->assign(
                (const char*)s, (size_t) key_size);
        resp->set_api_status(types::API_STATUS_OK);
    }
    else {
        resp->set_api_status(types::API_STATUS_ERR);
    }
    return ret;
}

hal_ret_t crypto_asym_api_fips_rsa_sig_verify(internal::CryptoApiRequest &req,
        internal::CryptoApiResponse *resp)
{
    hal_ret_t           ret = HAL_RET_OK;
    uint32_t            key_size;
    pd::pd_capri_barco_asym_fips_rsa_sig_verify_args_t args;
    pd::pd_func_args_t          pd_func_args = {0};

    key_size = req.fips_rsa_sig_verify().mod_n().size();

    switch (key_size) {
        case 128:   // 1K
        case 192:   // 1.5K
        case 256:   // 2K
        case 384:   // 3K
        case 512:   // 4K
            args.key_size = key_size;
            args.n = (uint8_t *)req.fips_rsa_sig_verify().mod_n().data();
            args.e = (uint8_t *)req.fips_rsa_sig_verify().e().data();
            args.msg = (uint8_t *)req.fips_rsa_sig_verify().msg().data();
            args.msg_len = req.fips_rsa_sig_verify().msg().size();
            args.s = (uint8_t *)req.fips_rsa_sig_verify().s().data();;
            args.hash_type = req.fips_rsa_sig_verify().hash_type();
            args.sig_scheme = req.fips_rsa_sig_verify().sig_scheme();

            args.async_en = false;
            pd_func_args.pd_capri_barco_asym_fips_rsa_sig_verify = &args;
            ret = pd::hal_pd_call(pd::PD_FUNC_ID_BARCO_ASYM_FIPS_RSA_SIG_VERIFY, &pd_func_args);
            break;
        default:
            HAL_TRACE_ERR("Unsupported key size: {}", key_size);
            ret = HAL_RET_INVALID_ARG;
            break;
    }
    if (ret == HAL_RET_OK) {
        resp->set_api_status(types::API_STATUS_OK);
    }
    else {
        resp->set_api_status(types::API_STATUS_ERR);
    }
    return ret;
}

hal_ret_t cryptoapi_invoke(internal::CryptoApiRequest &req,
        internal::CryptoApiResponse *resp)
{

    hal_ret_t           ret = HAL_RET_OK;

    switch (req.api_type()) {
        case internal::ASYMAPI_ECC_POINT_MUL_FP:
            ret = crypto_asym_api_ecc_point_mul(req, resp);
            break;
        case internal::ASYMAPI_ECDSA_SIG_GEN_FP:
            ret = crypto_asym_api_ecdsa_sig_gen(req, resp);
            break;
        case internal::ASYMAPI_ECDSA_SIG_VERIFY_FP:
            ret = crypto_asym_api_ecdsa_sig_verify(req, resp);
            break;
        case internal::ASYMAPI_RSA_ENCRYPT:
            ret = crypto_asym_api_rsa_encrypt(req, resp);
            break;
        case internal::ASYMAPI_RSA_DECRYPT:
            ret = crypto_asym_api_rsa_decrypt(req, resp);
            break;
        case internal::ASYMAPI_RSA_CRT_DECRYPT:
            ret = crypto_asym_api_rsa_crt_decrypt(req, resp);
            break;
        case internal::SYMMAPI_HASH_GENERATE:
            ret = crypto_symm_api_hash_request(req, resp, true);
            break;
        case internal::SYMMAPI_HASH_VERIFY:
            ret = crypto_symm_api_hash_request(req, resp, false);
            break;
        case internal::ASYMAPI_RSA_SIG_GEN:
            ret = crypto_asym_api_rsa_sig_gen(req, resp);
            break;
        case internal::ASYMAPI_RSA_SIG_VERIFY:
            ret = crypto_asym_api_rsa_sig_verify(req, resp);
            break;
        case internal::ASYMAPI_SETUP_PRIV_KEY:
            ret = crypto_asym_api_setup_priv_key(req, resp);
            break;
        case internal::ASYMAPI_SETUP_CERT:
            ret = crypto_asym_api_setup_cert(req, resp);
            break;
        case internal::ASYMAPI_SETUP_PRIV_KEY_EX:
            ret = crypto_asym_api_setup_priv_key_ex(req, resp);
            break;
        case internal::ASYMAPI_FIPS_RSA_SIG_GEN:
            ret = crypto_asym_api_fips_rsa_sig_gen(req, resp);
            break;
        case internal::ASYMAPI_FIPS_RSA_SIG_VERIFY:
            ret = crypto_asym_api_fips_rsa_sig_verify(req, resp);
            break;
        default:
            HAL_TRACE_ERR("Invalid API: {}", req.api_type());
    }
    return ret;
}

} /* hal */
