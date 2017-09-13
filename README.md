# PMd

Simple secure storage.

Public server available at https://pm.lessmore.pw

<!-- markdown-toc start - Don't edit this section. Run M-x markdown-toc-refresh-toc -->
**Table of Contents**

- [PMd](#pmd)
- [API](#api)
    - [User registartion](#user-registartion)
    - [Push data](#push-data)
    - [Get data](#get-data)
    - [List data versions](#list-data-versions)
    - [Delete data version](#delete-data-version)
    - [Selfdelete](#selfdelete)

<!-- markdown-toc end -->


# API

## User registartion

Send post request with you PGP public key.
```
PUBKEY=`gpg --export -a <you_keyid>`
cat > reg.json <<EOF
{
    "pubkey": "$PUBKEY"
}
EOF
curl -XPOST -d @reg.json https://pm.lessmore.pw
```

Return JSON:
```
{"error": "","userid": 100}
```

## Push data

Put you data in JSON and sign
```
SECRET=`cat my_secret | base64`
cat > data.json <<EOF
{
    "data": "$SECRET"
}
gpg --clearsign --sign-with $PUBKEY > signed

curl -XPOST -d "`cat signed`" https://pm.lessmore.pw/push
```

Return JSON with data version, time, and state
```
{"version":"CyVYq3do1DM1mMf96ziAGfZB3s856zH6eWQVoQsMzgQ=","time":"2017-09-13T16:51:53.025960496+03:00","state":"OK","error":""}
```

## Get data

Sign JSON request
```
echo '{"version": "CyVYq3do1DM1mMf96ziAGfZB3s856zH6eWQVoQsMzgQ="}' | gpg --clearsign --sign-with <you_keyid> > signed.json
curl -XPOST -d @signed.json https://pm.lessmore.pw/pull 
```

Return armored encrypted data
```
-----BEGIN PGP MESSAGE-----

hQEMA/jh+hN2m5ysAQf/X4eTuoIZyIghBWBMsfr72RgIcquv5uE6QnybfZH2Aq8A
f1rzhbBfbVlMgLwvKY/I7C9G5ZGCQQViyC7VaWxqvI8V3sOe8IEDyi1DqEzuxZA+
FQgIEFJpbWRo0Fu4/tQmUuDYiVIFHC2h2/jc8+Kj8KR12hWnvL5mjKogqsTSpOtk
OMCzcZs/1p298TlLx/RZoY+Ktn6IvRtE2PdEw0Kw7F4g9Jad3wyGykdH81rPETTt
GPtV7CmaIVE4sSdWgDA/fkR8Gin0cKXVxC+c54wXs+iRpgX/nE1ZFGaPNG5MG117
u/5S7mDg3S6Znaym38Mqu1oav7ImmsfvPNV8Q1LT+tJJAVqzJhDJY83XCTIDQict
5xw0c6PnfRPUffvmdzrstSU35LV+tMz9kDGqjW1Ss40VC3PKixg0yk6Jz5SamLBS
4sKoE8PDmL+qCw==
=2qKU
-----END PGP MESSAGE-----
```

## List data versions

Make POST request with some signed message
```
curl -XPOST -d @signed https://pm.lessmore.pw/list
```

Return armored encrypted data with JSON structure:
```
{
    "state": "ok",
    "versions": [
        {
            "time": "2017-09-13T13:40:58Z",
            "version": "pjqzXjNL_hIpKt_gcXXgKEIb5Caile2wug0cmCcTx-k="
        },
        {
            "time": "2017-09-13T13:41:10Z",
            "version": "n-bST3_-TtA1d-4S1upjBO3c5tGQeBgonYa5G-WvhMo="
        }
    ]
}
```

## Delete data version

Not implemented yet

## Selfdelete

Not implemented yet
