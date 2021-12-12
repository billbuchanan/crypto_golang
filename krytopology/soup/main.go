

	/*
		//
		var p, q *big.Int
		var bits int

		bits = 99

		p, _ = rand.Prime(rand.Reader, bits)
		q, _ = rand.Prime(rand.Reader, bits)
		group, _ := camshoup.NewPaillierGroupWithPrimes(p, q)

		domain := []byte(mymsg)

		ek, dk, _ := camshoup.NewKeys(1, group)


		msg := big.NewInt(1)
		cs, _ := ek.Encrypt(domain, []*big.Int{msg})

		cs, proof, _ := ek.EncryptAndProve(domain, []*big.Int{msg})

		err := ek.VerifyEncryptProof(domain, cs, proof)

		if err == nil {
			fmt.Printf("%+v\n", err)
		}

		dmsg, _ := dk.Decrypt(domain, cs)

		fmt.Printf("%v", msg.Cmp(dmsg[0]))
	*/
