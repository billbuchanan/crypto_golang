# crypto123
Secret shares and other things

Example command:

# Decrypting 3/4 secrets:
```
    ./main -decrypt -c'31a290f90be3d36d94d5a2f160b2273a41867fc08943b426f454d282e364cd7f737c51d610048ea9227aa1645a06670b7f8cb9ff9c' -s'dU8i-aASwHad2XlRMthylWosA2aNrMUsQEaC-4wFQes=d3A1O38NSGd3s_3X4Y_J1H6sPP-fZTCH5uf_j1lVUKE=' -s'Vn9pJ8GyjdQoMf25-G1IC0aVAhyL1oK3rV5LMJZrhnE=ckEzlhMlGkMBEDYKgc_qUT_FlVJvSSpSii-3e3slrnI=' -s'WRegt-t227E5W7WGjMWNUISTekU47Okv9OP6UMC7YS8=-uScKbV4iIsp5LtqIy2moHPYR7xTdBKvh4l0hCMs7r4='
```
Expected output:
```
    plaintext: Test3
```
# Generating 6 secret shares of which 5 are required to combine:
```
    ./main -encrypt -s'Test' -n6 -k5 
```
Expected output (will not be exactly this):
```
    Ciphertext:'be8b3f6d6071b07285cc264b3f38d76a19b2a35b642c550f5d8c61a9cd5d2370ffc6460c8ad86737c0b1be8be6b4ff85bc6d2656'
    Shares:
    'OuJRNYcUXMmGVQucvraunNRliiK8Wy19Em0WeYp3sac=UWzOnFQs3IL9Y_4BNBH19d2ywnLp2sBTek8h1MHsaYQ=',
    'DcL8X_DgomN5XGYNS2teZWXoU3LOiiq52zeGxi3SMxg=W6KvKR_vI9Fw6qeOO5o-Og8u3mDTwiQLh7d1DLnKBM4=',
    'XlYoGb8JErLC1NzP2ntIRFCYgaBBf0w8SHrKd5gGTgU=aNwOHjuTQj2rFPvIT3Om-qj5uDct4pZUHBEF8afU3zc=',
    'uEi7MATJyOa-BtZD3AYvzNN-IVD4fIiKDtoXJbpKK-I=kEPY5zs7xrxUaoQ3RK24fkraNsmA7afVoIXTP4VFVwI=',
    '4tiTAAp8OPb3VOThzYLnqLvkGAJKDBAMeprnaxJAIqE=XlOcJNG_JGxsn0WOdIs3O2g8KYtIjcqBJObvdvhfDok=',
    'LXVBEdm_lXGAQurlYtuvx-9z-29Azb9JCzeGmQUn8hU=K7JTGiP5s35_Pn9r67UGKP4y2-xxYl6iSOwCnWdxHDo=',
```

# Output 4 secrets to a folder with 1 file per share:
```
    ./main -encrypt -s'Test' -n4 -k2 -foldertest
```
Expected output
```
    #test/0
    77b5346f688ba67704b5d028646ebcd59870542b58a32da12a7fc8ebf4000f73abc3dc40b83fc7bd0b7013ad5df8e7ec0b8eb2d4c6,KDE0hif2am8XA5re9r7zwNdwTBz-P12R4UpLnl2vCSw=Kz0UlvMw_VaCLyoA191kuo3Rh23RK7U35j38-RVEjAs=7eb27d507d72c4d1fb11ff6612d7a413436d1e97a4b769bf9d13496b35134f016c8553e46cb81bb9c667794c2ca95246fb1724572b,01bwfQu94-to2dgKRF2kSE4nOvECikqL52GlIFFOssI=zAgsCMqHiSOSpBcm9F9DAimKNRWU3FPpTnsbZ7-bRz4=

    #test/1
    77b5346f688ba67704b5d028646ebcd59870542b58a32da12a7fc8ebf4000f73abc3dc40b83fc7bd0b7013ad5df8e7ec0b8eb2d4c6,cXXdWi9QSrtB8eDIhaD0NHDvKnKmZ3h1QhuWI0usYnU=Kz0UlvMw_VaCLyoA191kuo3Rh23RK7U35j38-RVEjAs=7eb27d507d72c4d1fb11ff6612d7a413436d1e97a4b769bf9d13496b35134f016c8553e46cb81bb9c667794c2ca95246fb1724572b,TinxUsNu1bDod_4s_0LsDVC9Zrb2N6v07A5CrJGKEak=zNYwOkJOMOZdjZ1ln3QqkgnPnIGuomgyciQtbSPAkXU=

    #test/2
    7eb27d507d72c4d1fb11ff6612d7a413436d1e97a4b769bf9d13496b35134f016c8553e46cb81bb9c667794c2ca95246fb1724572b,iet2kX9y0WT4p5BFZBB6ixDch442YELQmAAiTwPXyNI=iYlpL513H6P0Sp6aT77op6JQkot7uwQC3vAyjt6RlZc=

    #test/3
    7eb27d507d72c4d1fb11ff6612d7a413436d1e97a4b769bf9d13496b35134f016c8553e46cb81bb9c667794c2ca95246fb1724572b,vNnny5njAewLptqw4CiuRVMQ1U8EVAprruS9fzhsAH8=QobBDJ__J7CtdzkZBuHHKWO6Kt4tld21ke2daMGdKxk=
```

# Take the input from a folder of shares (must be the only thing in the folder right now):
```
    ./main -decrypt  -foldertest
```
Expected output:
```
    plaintext: Test3
```


# Output to csv - will be appended if the file exists
```
    ./main -encrypt -s'Test' -n4 -k2 -csvtest.csv
```
Expected output: 
```
    #test.csv
    4813d17bf9d113e0b628941aafbd1fa1ae87ca63e766a97e18368fcca47a48edc840011a504b522c99d72ed3c38145f271cd3882bc,QSG-i98DFz5M7AfAJ1B0vsSIo-yFE6E5qoiaLF0UpLM=KkSUfmaGBRRE2Fe_313Ra8O7oAGYcrlLRdPPYASF4Hc=,BUMSNhhl_5S37p6i3tsbP5FvZyoz5aIo9Y_8flclX0Y=eMRt1aJ6ldESQZ3gKjwY8YWN8JOgMtmk8rdVEbR7HEc=,dx6UCncs5_BiWuZNChGsWxRcP1J2E4Q4Km-EZN03H-Q=_2kl3xjyV0LvrDHADf8mB78NHyX6kWiJ19FFmqoCCGw=,6BP03vbR8siaMyf6_yArnFnyFhIHT4s1d24xNExK5ZY=Qwi4HEBTYKnJPA0x2ZijiuhatqN74NPqTlW11rHLdDo

```

# Speed test 
Runs on a dataset of 100,000,0 with three libraries:

    github.com/codahale/sss
    github.com/dsprenkels/sss-go
    github.com/SSSaaS/sssa-golang

```
    ./main -speedtest1000000
```
Expected output:
```
    Doing 100000 of tests with 5 keys which require 4 to recombine

    Codahale test on private keys:
    Creating 500000 shares took: 10.831973332s
    Combining 500000 shares for 100000 secrets took: 1.363223701s
    Codahale took: 12.195353448s



    Dsprenkels test on private key data:
    Creating 500000 shares took: 457.080476ms
    Combining 500000 shares for 100000 secrets took: 596.890736ms
    Dsprenkels took: 1.054213302s


    SSSaaS test on private keys:
    Creating 500000 shares took: 10.029006ms
    Combining 500000 shares for 100000 secrets took: 162.697889ms
    SSSaaS took: 172.923588ms
```

# PHP frontend (proof of concept, very insecure right now)
```
    php -S localhost:9000
    connect to localhost:9000/
    Type secret into box, press submit secret
    You will get a zip file downloaded.
    Unzip this, then drag the shares into the dashed box.
    Press decrypt. 

    Press Reset to delete the current uploaded files on the right. You will need to delete these to upload a new shares, as otherwise the output will fail.
```