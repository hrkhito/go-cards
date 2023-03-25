// mainは実用的である特別なパッケージ
package main

// type num []int

// func (n num) roop() {
// 	for _, num := range n {
// 		if num%2 == 0 {
// 			fmt.Printf("%v is even\n", num)
// 		} else {
// 			fmt.Printf("%v is odd\n", num)
// 		}
// 	}
// }

// main関数はmainパッケージをビルドして作られた関数が実行されるたびに自動的に呼び出される
func main() {
	// var card string
	// card = "Ace of Spades"

	// var card string = "Ace of Spades"

	// card := "Ace of Spades"
	// card = "Five of Spades"

	// card := newCard()

	// fmt.Println(card)

	// cards := []string{"Ace of Diamonds", newCard()}
	// cards := newDeck()
	// apends関数は既存の関数を変更するものではなくて、新しい関数を作成する
	// cards = append(cards, "Six of Spades")

	// スライス内の全てのレコードを繰り返し処理するときに使用するのがrangeである
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	// cards.print()

	// hand, remainingDeck := deal(cards, 5)

	// hand.print()
	// remainingDeck.print()

	// greeting := "Hi there!"
	// fmt.Println([]byte(greeting))

	// cards := newDeck()
	// cards.saveToFile("my_cards")

	// cards := newDeckFromFile("my_cards")
	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()

	// nums := num{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// nums.roop()

	// for i := 0; i < 11; i++ {
	// 	if i%2 == 0 {
	// 		fmt.Printf("%v is even\n", i)
	// 	} else {
	// 		fmt.Printf("%v is odd\n", i)
	// 	}
	// }
}

// 関数が値を返す場合は戻り値の型を定義する必要がある
// func newCard() string {
// 	return "Five of Diamonds"
// }
