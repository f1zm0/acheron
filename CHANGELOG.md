# Changelog

## 1.0.0 (2023-04-24)


### ⚠ BREAKING CHANGES

* change signature for syscall func
* remove direct syscall support
* replace djb2 with xored version in asm
* change hasher interface to func type
* done resolver logic and add gadget search in asm

### Features

* add hash func helper in package entrypoint ([81bc1f4](https://github.com/f1zm0/acheron/commit/81bc1f4698329e62822a5bc0d539aeb119979fc2))
* add helpers for errors ([4d57328](https://github.com/f1zm0/acheron/commit/4d573287d18a9be0a41d3f917df06be659783931))
* add internal resolver and util packages ([f3c3edb](https://github.com/f1zm0/acheron/commit/f3c3edb5fb6b07710d0c416197fcaad623033b1a))
* add lib entrypoint and indirect syscall asm ([7db263d](https://github.com/f1zm0/acheron/commit/7db263dd8e84094c37317457bf5e66fce179c551))
* add nosplit flag to asm routines ([f995154](https://github.com/f1zm0/acheron/commit/f995154072f775089d4583ddd9e8608299d18e1f))
* add process snapshot example using acheron ([5ac8f2a](https://github.com/f1zm0/acheron/commit/5ac8f2a4fc0d68d7fc247bcf0c160b4e80059829))
* add zw exports parsing logic to resolver ([19f45c7](https://github.com/f1zm0/acheron/commit/19f45c70eb4eb23c5a0bcba96f2e53e12190bae3))
* change signature for syscall func ([89ce53c](https://github.com/f1zm0/acheron/commit/89ce53cead6765a1cbfa2fc9d746388736cdb550))
* done resolver logic and add gadget search in asm ([22a9c23](https://github.com/f1zm0/acheron/commit/22a9c23039bc727a8356e77bef38333335d15ef5))
* ported ntdll module parsing to go assembly ([21f66df](https://github.com/f1zm0/acheron/commit/21f66dfe52de024600c61d7b9340cc9a69f2d0be))
* replace djb2 with xored version in asm ([4c2cd88](https://github.com/f1zm0/acheron/commit/4c2cd8888575f367354ea5d8748958fb8d0cc5a3))


### Bug Fixes

* correct offset for return value in gadget search ([8de5eec](https://github.com/f1zm0/acheron/commit/8de5eecf01c5c1448ca2e84293179efc43b3915a))
* fix broken indirect syscall asm implementation ([8c9d99d](https://github.com/f1zm0/acheron/commit/8c9d99df0720e238c767178bcae95d351728b765))
* fix helper function names after renaming ([0783417](https://github.com/f1zm0/acheron/commit/0783417cb38b5998da028f8160a580b76c1d6621))
* fix inverse check in value comparison in error helper ([e040b8a](https://github.com/f1zm0/acheron/commit/e040b8aadf42e891dbaa478340d7458e010c4c35))
* update helper asm routine names ([6be7b5a](https://github.com/f1zm0/acheron/commit/6be7b5a8e7b8efb2283be5862e7a10ff013089a2))
* update names of asm routines ([fd67f62](https://github.com/f1zm0/acheron/commit/fd67f62086abbce8c992dbc6c483151c9fa82a22))


### Continuous Integration

* add release workflow ([f6305de](https://github.com/f1zm0/acheron/commit/f6305deb8247a842d35e81d31782765d74dfff84))


### Code Refactoring

* change hasher interface to func type ([9e2294c](https://github.com/f1zm0/acheron/commit/9e2294ca16715f432b7522f7ae6f5e5b8bbcd017))
* improve error handling resolver and syscall ([2d2fb24](https://github.com/f1zm0/acheron/commit/2d2fb24611811c355421335cf70baa212e44249d))
* remove direct syscall support ([919ad92](https://github.com/f1zm0/acheron/commit/919ad920eefb706fc33c650e7c24b1d735d7917f))


### Documentation

* add custom hash function example ([26844d5](https://github.com/f1zm0/acheron/commit/26844d5e49a47189c66eaae7dfb983b1b16ddcc0))
* add example for direct vs indirect syscall comparison ([8060bbf](https://github.com/f1zm0/acheron/commit/8060bbf9299e358741ce7fc0baf5872b39fd4b78))
* add examples summary table ([ac7afc2](https://github.com/f1zm0/acheron/commit/ac7afc2e8d35db4686b75ab7d883170a986e32fa))
* update examples in readme ([db97e78](https://github.com/f1zm0/acheron/commit/db97e78a085660e09eb022395a69b2840eceadd6))
* update examples readme ([c381280](https://github.com/f1zm0/acheron/commit/c38128028cbbbbae6317cb28b7342dedd09c0940))
* update examples to reflect api changes ([f85a9d0](https://github.com/f1zm0/acheron/commit/f85a9d06873219a3b96e3fab442219ef7335ea63))
* update main readme ([5d41837](https://github.com/f1zm0/acheron/commit/5d418371e4be7796c7aad3999b0f61d389627d3c))
* update process snapshot readme ([a3a897f](https://github.com/f1zm0/acheron/commit/a3a897f2faff224ba41b9a0dae3096e37205dc9b))
* update readme for proc snapshot example ([2f6206a](https://github.com/f1zm0/acheron/commit/2f6206a33d25fbef52bdc3914397e185b514ff89))
* update sc_inject example ([d5fc72d](https://github.com/f1zm0/acheron/commit/d5fc72d3e763b811561bec30691d4048119eca83))


### Misc

* add comments to exported functions ([428e9ba](https://github.com/f1zm0/acheron/commit/428e9ba6a5ea0273babff63c2a7c3e9340ae5c4c))
* add exe to gitignore ([624e035](https://github.com/f1zm0/acheron/commit/624e0359d207dc9f9bf7e318b64486a8b33ec9a1))
* add gitattributes file ([f6ded30](https://github.com/f1zm0/acheron/commit/f6ded30da3228604652885a5f8ac5797c666fac2))
* add info and fix markdown syntax in readme ([d36180e](https://github.com/f1zm0/acheron/commit/d36180efc1bdc955711e25d72f217ae3d88bd900))
* add notes to syscall asm for ret code ([6da875b](https://github.com/f1zm0/acheron/commit/6da875ba8d0398ebea26adcaab56724cdb96e5a1))
* add package info to readme ([3839328](https://github.com/f1zm0/acheron/commit/3839328922aacca82bd12f21a535eaa0b6d40f7d))
* add readme banner ([98ab801](https://github.com/f1zm0/acheron/commit/98ab8018f73974c71f77d897efc0480eabf0c566))
* change hash func return type ([011fe8d](https://github.com/f1zm0/acheron/commit/011fe8d32528845a91414020c8b5e6c10d27895d))
* change instance var name for better distinction from pkg name ([887c346](https://github.com/f1zm0/acheron/commit/887c346afd50428f61879390c20937e6a4c8fefd))
* change slice var name for clarity ([5f1b1ce](https://github.com/f1zm0/acheron/commit/5f1b1cec3a837915dbdc8de45a4894300a8e8f44))
* cleanup and add comments ([68036a7](https://github.com/f1zm0/acheron/commit/68036a7c3197810b002d5f0a0bae83940643681e))
* fix incorrect newline in code example in main readme ([7b90910](https://github.com/f1zm0/acheron/commit/7b909109505a31be8512f4f9b5508400e8c3f6bd))
* fix wrong indentation ([4312610](https://github.com/f1zm0/acheron/commit/43126101a466d56c96cb60aebec55e2b7bfe13d7))
* fix wrong indentation in asm file ([a9d63de](https://github.com/f1zm0/acheron/commit/a9d63de1378050de5ba30776688161d6c63defde))
* fix wronge indents in asm file ([12e929d](https://github.com/f1zm0/acheron/commit/12e929d454747d46284f28701d3f820aa4d5d66c))
* improve comment in library entrypoint file ([432af39](https://github.com/f1zm0/acheron/commit/432af392dc44c51a8d562898212626c88efcaf34))
* initial commit ([b6a503d](https://github.com/f1zm0/acheron/commit/b6a503dca255d0a726945fbe316ddaee3b4c76e2))
* license change ([17ed9cf](https://github.com/f1zm0/acheron/commit/17ed9cf1359fe4ca134627948a58e8e5cc38fd17))
* move examples table to main readme ([4c80756](https://github.com/f1zm0/acheron/commit/4c807568c6557d0358893e08f3ee81e270a4ae65))
* move lib functions to entrypoint file ([88edc76](https://github.com/f1zm0/acheron/commit/88edc76abe22fe93c4fa15b3ebf0f0de25526573))
* move syscall routines and stubs ([6dc95b5](https://github.com/f1zm0/acheron/commit/6dc95b5a7992cbb976352ca3d845f249d17b1998))
* remove unused pkg ([23d1503](https://github.com/f1zm0/acheron/commit/23d15039780835a128268d7adb458f327f4bbbc4))
* remove unused types struct ([3ceb34f](https://github.com/f1zm0/acheron/commit/3ceb34f67ac328d0340d6bdb8ca580e826ebdb4b))
* rename asm file for consistency ([26b26d7](https://github.com/f1zm0/acheron/commit/26b26d74759f64d34a4603079652518a02e17876))
* rename asm files ([7b33816](https://github.com/f1zm0/acheron/commit/7b33816d3b28d28af937413197cbcb6a273c10c4))
* rename memory read functions ([859833e](https://github.com/f1zm0/acheron/commit/859833e20f3633d79d518112bafd3cfaa318ca43))
* rename vars and struct members for clarity ([25f21db](https://github.com/f1zm0/acheron/commit/25f21db656006471a272ce029fd5a4b340a2669b))
* update gitignore ([13923a5](https://github.com/f1zm0/acheron/commit/13923a509d5260778f0435972c68f151cab22637))
* update mod and sum files ([58a25b4](https://github.com/f1zm0/acheron/commit/58a25b46d5f30e4cda9c60efb601509e5b94111a))
* update mod and sum files ([4c0a173](https://github.com/f1zm0/acheron/commit/4c0a173151918f5bf3a6bf0f337c00d1f266ef12))
* update mod and sum files ([bc4c07d](https://github.com/f1zm0/acheron/commit/bc4c07df50981f01dcabc457606d0e3077a14dc6))
