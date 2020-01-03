package main

import (
	"github.com/markbates/pkger"
	"github.com/markbates/pkger/pkging/mem"
)

var _ = pkger.Apply(mem.UnmarshalEmbed([]byte(`1f8b08000000000000ffec595b8fdb3616fe2b015f57317597acb7cd069d06bb09da24ed4b1104144549cc50a44a5213bbc1fcf705a9ab3db6ec60b36d801a8331790e3fde0e0f3f1e8a5f00e5a55020fb022aaaeb2edf60d14055239ea33dac448ef6a6ec25952003f01745a4829f4859c2568a4f43b9035e35ad90fa27a46b909d6fc7016f50430c62105f0a0c32001cf01ec98ae83eff5608bdd6d76ba4710db2dfc0067c70c03b8d180199961d1984b70429c1410694919e15a425bc201cefb3678b91958a0b4dcbfd94010eb8133f504694691bb5eda612a6837e6a5699ef3551c00165a3817330cdb6f402f8405b228103a88054749a32e00061f02dd2352c292326031ca0b4a4bc32256acfb1e9e42569fb1eba920ae04c3d61d1b49228054b8634592aaa3f686b65ae11e5444246951e1464677372df6a316520ea5bec054cdbda0e76908b6561a1d02c107c28167e1479db270a48b92692230649f119c9421dc318a3ada678d6d40d5a48537589783118efb84875b966642e688a68164cbd8584c385b09c80aa917720f9517c20479ebf908fbad46c61a75de46e0f25d8ded31d7000e1581494578b2c448a7b4b39478ac4e181867224f74b0d560f4bb126cbc6e127e3e70bb9258d11a514d2ba2943d5296f3de5fa8be21aa99a62215b5863b65264ff91d297202d92cabada3a4ada65be045318717e054c4b2cf8c3259816f784af818c85af9880855d31348b3bd169832a8a0547541aaa6b89d4d4eec725846a5c13c66ad8a05669d961dd49728869096344532261259e6bd1b0532c854a22c5d902d8581f7a52889f2c74affff419314de46782744d6483f829503b78e259c65c1674b950540b58094dfa15140cf16a2364057750ed15ecb8dd63077ac379504bc455296473b2b4e3148b82403e02dafb6a4339a49c6e1ebca5668f1ab679302c60966f48209638b03a6d0d6b12a849d30ecc3cb194e16ec4aaa50ab7dd522c1bad84d44b15275a4b840f1a12ca6ee5a5aa158c2de5e32a92948c60cde8e19014e5152325a3557dd0abda2b8c181b2d3ae935519a093b057174a0f5eaa63fc84c02735acd59adc6fcc0c70d6dc890c0a6639ab6c8cedc2a7eef842685ddfa28b734cb89ee7f61ad75bbc8da9fd14493721cf1a033ebdc4a614f3d2377723a8385826447f0701a9f389407d3d95c4576ed94816acf3532f6911dd7fd74861cc495584893fd90168d3dea9e940c867ba2577b1b0ef45e31b3d751a4d02773f3c3fa0107f4fe8ff2cf92f62cda2b66f73c907b4a33da7e3ac3d69873b0d3a5171fcaa915152a0dee81f042189e59ecb231b4a811ae91efb682edbdc08d2ea06d62f6ccb5b8f1205d0377f2818c01c80aaebe2fca75c4bc8587595d099fe28615f0050319072eb832ff0d510a55e79a3bd8225567b7e0455c2bc56e7f01e8c3ba45f87e05450b8e568aa5e8ceaeab61f29e174f955a7755047792c09c16547667adf984facf824667360d5e83b327c50707bc274a1fde0d3e1a861c2e08a674714998c2af83f35d6b6e0e65f5bba1e6e0e8c8d392685c4b4bbb262e434a11797cdc1e831a6157c68abc3203792d0a63a5ec0bb87c117b8d281f6f4d17ee7677e2b528ce026025368d282cee572215b5572f6fe305e0f1f1d101656fb5f3b7cb0ce2a618da1af9c95630f75293164423caac8af7d7c719e60045ff2020f3fcd4018da1b0ccf7c2244c432ff4ace6a3a5b80cf8aeef3e77bde76ef0de8bb230cebc78e3265192c6a997fec34d33d76c6daa3e16c61abd610c299bbb2f790099172749e2fbae35b3e9314e93c4016f18e5f7200b1df08a0b9085419824811f27ae037ea105c8223770c09dc999aa6f0bd394eb809f50f11157e2a30bb2df5cc7fe7d70c03fcd58554bb0e9f59d49bc2849dd208992d8016f94d1982127ee367c74c0eb2be166868f0ef8d7d7c15f50a9eb3355e26d104c5552dff3dcc88d1f1df06e5e8c174ce07b6527fb82ddf7ab14badbd8013f3054f5057784dbf43faa4592d8eccf43d658e5837520734dfd4e3f409ced67fdbb44de51563c7bf5f259435563dbbaf4e981778c8d1f09fafc5750525fe1fb238721123323baafae5ab5934cf1e880026934ced3b80fd773ab736ddbe557b11044eb3c8416fc33d24f68f6f959de89c22c8a365e94865ebcf5fc27bc5322a62e104f1078693a118f77403ca11b24df927852779b8cfb3c48d2c8f582648d7952771b8df869922bd4730ebfc63d499478d398d2348e1237f016dc33534ffa2750cf49973ca2a3d905979035ea9949a677ea814c06ef3864932579f4e86b696356fd4db802a2992d00bafb55e5776955f8db7d7ef7f3bfc137e68e7c9d3bf2893b8c635e1dbaa499bfcddc681384491487dbd8fb9f4397e08841c26fc920611244d36e0d93681ba4e13638cf2061128413234c733ccf2067f117186419f0c45e6c039e77d36adc62975becf2a7f051fe7f8d5e72b8dbedd659c800461e8aaf0961d22c0cb320d90451b24de3244dbe710813267ebafd960494fa493a6ef66de426a9b775dd15024afd640e79c649ae10d039fc0a013d21ad304ac29981e2bf3e8281f9f91826bfc52e7f0d57d8dd3cc72fe4c73a243f5617e396de0bc8f77a0edd1ef16f8ff8b747fcdb23feed11fff6887f7bc4bf3de2df1ef16f8ff8b747fcdb23feed11ff7bb9bf3efe170000ffff010000ffff43b65b56232e0000`)))
