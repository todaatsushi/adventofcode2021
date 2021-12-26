const puzzleInput: string = `
  176
  184
  188
  142
  151
  156
  157
  167
  166
  178
  182
  191
  190
  191
  192
  196
  197
  201
  204
  207
  212
  213
  231
  232
  234
  232
  239
  268
  279
  280
  282
  278
  283
  284
  283
  296
  302
  320
  321
  332
  334
  332
  335
  340
  347
  348
  352
  342
  334
  333
  332
  357
  365
  366
  386
  394
  391
  387
  388
  389
  391
  394
  381
  392
  394
  397
  407
  428
  441
  444
  447
  448
  449
  450
  449
  453
  467
  476
  472
  486
  512
  517
  519
  530
  533
  537
  538
  552
  554
  550
  575
  583
  585
  597
  598
  600
  601
  602
  606
  607
  609
  611
  603
  604
  597
  598
  599
  606
  612
  615
  626
  621
  623
  624
  627
  630
  642
  641
  644
  642
  645
  648
  640
  636
  640
  660
  669
  676
  688
  689
  712
  716
  711
  715
  698
  705
  729
  730
  731
  714
  720
  710
  723
  727
  730
  738
  739
  743
  753
  755
  763
  765
  761
  758
  759
  780
  803
  788
  789
  810
  827
  837
  849
  843
  853
  852
  853
  865
  866
  876
  878
  871
  874
  878
  884
  888
  892
  887
  893
  894
  905
  906
  887
  912
  905
  898
  901
  879
  882
  897
  920
  918
  911
  916
  922
  931
  933
  931
  936
  943
  945
  955
  956
  951
  957
  953
  960
  965
  973
  994
  1012
  1016
  1022
  1025
  1028
  1031
  1032
  1037
  1038
  1031
  1035
  1053
  1057
  1074
  1075
  1081
  1080
  1060
  1063
  1053
  1055
  1053
  1062
  1063
  1060
  1062
  1065
  1066
  1063
  1064
  1074
  1073
  1081
  1095
  1089
  1106
  1111
  1097
  1098
  1101
  1107
  1111
  1113
  1125
  1118
  1133
  1142
  1141
  1144
  1145
  1153
  1165
  1167
  1174
  1178
  1158
  1165
  1174
  1175
  1191
  1194
  1192
  1189
  1190
  1193
  1194
  1196
  1214
  1238
  1246
  1272
  1273
  1287
  1288
  1296
  1298
  1299
  1315
  1320
  1324
  1326
  1341
  1342
  1341
  1368
  1375
  1386
  1387
  1388
  1399
  1397
  1398
  1408
  1407
  1409
  1410
  1422
  1419
  1423
  1419
  1415
  1416
  1446
  1468
  1479
  1484
  1488
  1489
  1491
  1495
  1504
  1505
  1504
  1520
  1540
  1547
  1541
  1548
  1539
  1541
  1548
  1544
  1545
  1547
  1532
  1552
  1566
  1570
  1571
  1569
  1571
  1572
  1543
  1544
  1546
  1547
  1576
  1574
  1592
  1597
  1600
  1608
  1609
  1623
  1631
  1647
  1640
  1657
  1660
  1677
  1674
  1669
  1670
  1654
  1667
  1680
  1696
  1698
  1699
  1700
  1706
  1696
  1697
  1699
  1698
  1699
  1711
  1715
  1689
  1690
  1701
  1705
  1739
  1741
  1745
  1737
  1729
  1749
  1751
  1758
  1759
  1756
  1758
  1772
  1773
  1783
  1807
  1806
  1813
  1827
  1831
  1832
  1835
  1840
  1857
  1866
  1872
  1875
  1874
  1879
  1890
  1883
  1905
  1908
  1922
  1946
  1936
  1939
  1941
  1945
  1950
  1951
  1956
  1950
  1951
  1952
  1957
  1958
  1969
  1972
  2004
  2013
  2026
  2028
  2034
  2043
  2048
  2064
  2074
  2068
  2081
  2063
  2064
  2065
  2069
  2047
  2046
  2033
  2035
  2034
  2035
  2037
  2038
  2033
  2038
  2039
  2046
  2047
  2048
  2043
  2083
  2086
  2093
  2107
  2108
  2091
  2105
  2113
  2109
  2106
  2116
  2119
  2126
  2127
  2132
  2135
  2147
  2148
  2143
  2150
  2151
  2153
  2166
  2171
  2173
  2175
  2178
  2180
  2192
  2227
  2231
  2253
  2255
  2256
  2257
  2258
  2259
  2269
  2277
  2279
  2282
  2308
  2309
  2333
  2340
  2346
  2354
  2355
  2371
  2372
  2375
  2393
  2394
  2400
  2404
  2400
  2402
  2393
  2388
  2403
  2402
  2422
  2424
  2417
  2416
  2421
  2422
  2431
  2428
  2409
  2405
  2415
  2411
  2413
  2415
  2419
  2423
  2431
  2432
  2450
  2460
  2466
  2458
  2461
  2462
  2463
  2471
  2476
  2489
  2490
  2491
  2492
  2490
  2491
  2495
  2511
  2514
  2516
  2514
  2516
  2519
  2525
  2527
  2501
  2503
  2506
  2513
  2537
  2543
  2547
  2554
  2570
  2580
  2601
  2611
  2609
  2611
  2617
  2620
  2623
  2615
  2619
  2633
  2648
  2649
  2662
  2670
  2671
  2675
  2707
  2708
  2717
  2718
  2719
  2739
  2742
  2737
  2738
  2739
  2759
  2771
  2783
  2784
  2800
  2805
  2820
  2815
  2836
  2837
  2839
  2867
  2869
  2868
  2874
  2881
  2883
  2909
  2903
  2904
  2909
  2912
  2914
  2920
  2921
  2923
  2928
  2929
  2930
  2909
  2915
  2927
  2928
  2930
  2923
  2921
  2918
  2911
  2915
  2920
  2929
  2930
  2940
  2942
  2951
  2969
  2980
  2985
  2989
  2983
  2986
  2994
  3002
  3017
  3027
  3042
  3048
  3050
  3048
  3049
  3073
  3080
  3082
  3071
  3094
  3124
  3132
  3133
  3137
  3138
  3142
  3161
  3176
  3188
  3185
  3189
  3193
  3194
  3195
  3196
  3200
  3211
  3212
  3215
  3223
  3224
  3230
  3234
  3227
  3231
  3218
  3213
  3218
  3239
  3225
  3216
  3224
  3226
  3229
  3230
  3231
  3212
  3225
  3234
  3237
  3236
  3237
  3243
  3252
  3257
  3261
  3260
  3263
  3265
  3266
  3272
  3279
  3280
  3285
  3284
  3290
  3291
  3290
  3307
  3312
  3342
  3344
  3343
  3344
  3343
  3342
  3343
  3340
  3344
  3348
  3349
  3337
  3346
  3357
  3361
  3369
  3370
  3383
  3375
  3371
  3375
  3397
  3401
  3407
  3408
  3409
  3391
  3394
  3393
  3396
  3399
  3404
  3407
  3409
  3411
  3412
  3415
  3433
  3437
  3449
  3446
  3447
  3448
  3444
  3445
  3441
  3456
  3479
  3488
  3487
  3491
  3496
  3485
  3498
  3500
  3533
  3537
  3544
  3545
  3544
  3553
  3559
  3567
  3590
  3577
  3586
  3584
  3594
  3595
  3594
  3603
  3617
  3619
  3624
  3652
  3627
  3622
  3647
  3648
  3650
  3667
  3668
  3664
  3672
  3678
  3688
  3689
  3696
  3704
  3690
  3691
  3692
  3683
  3685
  3684
  3689
  3693
  3694
  3693
  3696
  3704
  3717
  3722
  3718
  3690
  3691
  3687
  3688
  3685
  3687
  3688
  3701
  3698
  3699
  3700
  3702
  3708
  3704
  3711
  3714
  3719
  3730
  3731
  3730
  3751
  3757
  3760
  3766
  3767
  3764
  3761
  3742
  3771
  3789
  3810
  3807
  3787
  3792
  3791
  3792
  3797
  3802
  3801
  3809
  3817
  3850
  3851
  3853
  3860
  3861
  3866
  3862
  3853
  3857
  3861
  3868
  3869
  3870
  3879
  3893
  3894
  3903
  3912
  3944
  3955
  3956
  3960
  3986
  3969
  3970
  3974
  3976
  4006
  4008
  4013
  4038
  4043
  4056
  4071
  4076
  4097
  4098
  4099
  4109
  4111
  4112
  4118
  4126
  4128
  4129
  4125
  4124
  4126
  4127
  4130
  4150
  4153
  4168
  4187
  4186
  4171
  4166
  4172
  4182
  4201
  4226
  4227
  4251
  4250
  4263
  4294
  4317
  4318
  4332
  4330
  4308
  4309
  4310
  4313
  4329
  4331
  4349
  4350
  4348
  4345
  4347
  4342
  4340
  4342
  4359
  4371
  4387
  4412
  4411
  4412
  4415
  4422
  4423
  4427
  4455
  4456
  4460
  4471
  4485
  4491
  4492
  4491
  4492
  4521
  4527
  4532
  4539
  4552
  4537
  4551
  4562
  4566
  4581
  4588
  4590
  4591
  4590
  4595
  4596
  4597
  4598
  4618
  4617
  4618
  4621
  4622
  4628
  4632
  4628
  4644
  4629
  4643
  4635
  4639
  4640
  4645
  4647
  4639
  4641
  4642
  4645
  4657
  4645
  4650
  4655
  4671
  4674
  4675
  4670
  4671
  4702
  4715
  4714
  4722
  4723
  4727
  4737
  4733
  4761
  4762
  4763
  4766
  4772
  4773
  4774
  4775
  4778
  4779
  4786
  4787
  4788
  4792
  4796
  4797
  4795
  4796
  4816
  4807
  4839
  4841
  4842
  4843
  4860
  4862
  4839
  4842
  4843
  4844
  4846
  4848
  4846
  4844
  4847
  4849
  4855
  4859
  4852
  4855
  4864
  4865
  4889
  4892
  4891
  4893
  4890
  4891
  4892
  4914
  4913
  4912
  4915
  4918
  4920
  4924
  4927
  4922
  4914
  4913
  4912
  4920
  4913
  4915
`
const puzzleInputs: number[] = puzzleInput.split("\n").map((num: string) => parseInt(num)).filter((num: number) => num !== NaN)

const testInput: string = `
  199
  200
  208
  210
  200
  207
  240
  269
  260
  263
`
const testInputs: number[] = testInput.split("\n").map((num: string) => parseInt(num)).filter((num: number) => num !== NaN)

export { testInputs, puzzleInputs }
