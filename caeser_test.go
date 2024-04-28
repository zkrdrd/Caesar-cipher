package caesarcipher_test

import (
	"caesarcipher"
	"fmt"
	"sync"
	"testing"
)

func TestEncodeDecode(t *testing.T) {
	t.Parallel()
	t.Run(`test find min number`, func(t *testing.T) {
		testTable := []struct {
			Smeshenie int32
			Values    string
			Result    string
		}{
			{
				Smeshenie: int32(5),
				Values:    string("The city was founded in 1221."),
				Result:    "Ymj hnyd bfx ktzsiji ns 1221.",
			},
			{
				Smeshenie: int32(13),
				Values:    string("People from all over the country came to Nizhny Novgorod to buy and sell goods."),
				Result:    "Crbcyr sebz nyy bire gur pbhagel pnzr gb Avmual Abitbebq gb ohl naq fryy tbbqf.",
			},
			{
				Smeshenie: int32(24),
				Values:    string("Also, Nizhny Novgorod has an advantageous geographical location because it is located between two major cities — Moscow and Kazan."),
				Result:    "Yjqm, Lgxflw Lmtempmb fyq yl ybtylryecmsq ecmepynfgayj jmayrgml zcaysqc gr gq jmayrcb zcruccl rum kyhmp agrgcq — Kmqamu ylb Iyxyl.",
			},
			{
				Smeshenie: int32(8),
				Values:    string("In conclusion, I would like to say that despite all the disadvantages I like my city because I was born there and have a lot of good memories of it."),
				Result:    "Qv kwvktcaqwv, Q ewctl tqsm bw aig bpib lmaxqbm itt bpm lqaildivbioma Q tqsm ug kqbg jmkicam Q eia jwzv bpmzm ivl pidm i twb wn owwl umuwzqma wn qb.",
			},
			{
				Smeshenie: int32(5),
				Values:    string("Однако, есть и некоторые недостатки."),
				Result:    "Уйтепу, кцчб н ткпучухак ткйуцчечпн.",
			},
			{
				Smeshenie: int32(13),
				Values:    string("Вы можете ездить на автобусе, троллейбусе, можете даже поехать на метро."),
				Result:    "Пи щыутят тфсхяй ън нпяыоают, яэышштцоают, щыутят снут ьытвняй ън щтяэы.",
			},
			{
				Smeshenie: int32(24),
				Values:    string("Сейчас его население — 1 500 000 граждан. Это шестой по величине город в России."),
				Result:    "Йэбпшй эыж ешйэгэеаэ — 1 500 000 ыишюьше. Хкж рэйкжб зж ъэгапаеэ ыжижь ъ Ижййаа.",
			},
			{
				Smeshenie: int32(8),
				Values:    string("town — это маленький городок, а city — большой город"),
				Result:    "bwev — еъц фиунхдтрс лцшцмцт, и kqbg — йцудацс лцшцм",
			},
		}

		wg := sync.WaitGroup{}
		for _, expect := range testTable {

			expect := expect
			wg.Add(1)

			go func() {
				defer wg.Done()

				// Проверяем поиск наименьшего числа
				result := caesarcipher.Encoding(expect.Smeshenie, expect.Values)
				if expect.Result != result {
					t.Error(fmt.Errorf(`result %s != %s`, expect.Result, result))
				}
			}()
		}

		wg.Wait()
	})

	t.Run(`test find max number`, func(t *testing.T) {
		testTable := []struct {
			Smeshenie int32
			Values    string
			Result    string
		}{
			{
				Smeshenie: int32(5),
				Values:    string("Ymj hnyd bfx ktzsiji ns 1221."),
				Result:    "The city was founded in 1221.",
			},
			{
				Smeshenie: int32(13),
				Values:    string("Crbcyr sebz nyy bire gur pbhagel pnzr gb Avmual Abitbebq gb ohl naq fryy tbbqf."),
				Result:    "People from all over the country came to Nizhny Novgorod to buy and sell goods.",
			},
			{
				Smeshenie: int32(24),
				Values:    string("Yjqm, Lgxflw Lmtempmb fyq yl ybtylryecmsq ecmepynfgayj jmayrgml zcaysqc gr gq jmayrcb zcruccl rum kyhmp agrgcq — Kmqamu ylb Iyxyl."),
				Result:    "Also, Nizhny Novgorod has an advantageous geographical location because it is located between two major cities — Moscow and Kazan.",
			},
			{
				Smeshenie: int32(8),
				Values:    string("Qv kwvktcaqwv, Q ewctl tqsm bw aig bpib lmaxqbm itt bpm lqaildivbioma Q tqsm ug kqbg jmkicam Q eia jwzv bpmzm ivl pidm i twb wn owwl umuwzqma wn qb."),
				Result:    "In conclusion, I would like to say that despite all the disadvantages I like my city because I was born there and have a lot of good memories of it.",
			},
			{
				Smeshenie: int32(5),
				Values:    string("Уйтепу, кцчб н ткпучухак ткйуцчечпн."),
				Result:    "Однако, есть и некоторые недостатки.",
			},
			{
				Smeshenie: int32(13),
				Values:    string("Пи щыутят тфсхяй ън нпяыоают, яэышштцоают, щыутят снут ьытвняй ън щтяэы."),
				Result:    "Вы можете ездить на автобусе, троллейбусе, можете даже поехать на метро.",
			},
			{
				Smeshenie: int32(24),
				Values:    string("Йэбпшй эыж ешйэгэеаэ — 1 500 000 ыишюьше. Хкж рэйкжб зж ъэгапаеэ ыжижь ъ Ижййаа."),
				Result:    "Сейчас его население — 1 500 000 граждан. Это шестой по величине город в России.",
			},
			{
				Smeshenie: int32(8),
				Values:    string("bwev — еъц фиунхдтрс лцшцмцт, и kqbg — йцудацс лцшцм"),
				Result:    "town — это маленький городок, а city — большой город",
			},
		}

		wg := sync.WaitGroup{}
		for _, expect := range testTable {

			expect := expect
			wg.Add(1)

			go func() {
				defer wg.Done()

				// Проверяем поиск наибольшего числа
				result := caesarcipher.Decoding(expect.Smeshenie, expect.Values)
				if expect.Result != result {
					t.Error(fmt.Errorf(`result %s != %s`, expect.Result, result))
				}
			}()
		}
		wg.Wait()

	})
	/*
		t.Run(`test find max number`, func(t *testing.T) {
			testTable := []struct {
				Values []int64
				Result int64
			}{
				{
					Values: []int64{-700, 0, 1, 10, 1, 2, 3, 4, 5, 6},
					Result: 10,
				},
				{
					Values: []int64{300, 0, 2, -100, 1, 3, 4, 5, 6, 7},
					Result: 300,
				},
				{
					Values: []int64{0, 1, 2, 3, 2, 1, 2, 3, 1, 2},
					Result: 3,
				},
				{
					Values: []int64{90, 98, 78, 105},
					Result: 105,
				},
			}

			wg := sync.WaitGroup{}
			for _, expect := range testTable {

				expect := expect
				wg.Add(1)

				go func() {
					defer wg.Done()

					// Проверяем поиск наибольшего числа
					result := caeser.Finder(expect.Values...)
					if expect.Result != result {
						t.Error(fmt.Errorf(`result %d != %d`, expect.Result, result))
					}
				}()
			}
			wg.Wait()

		})*/

}
