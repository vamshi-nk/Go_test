package main
import "testing"

func TestHelloFail(t *testing.T){
	t.Run("Saying hello to people",func(t *testing.T){
		got := hello("vamshi")
		want := "Hello, vamshi"
		correctMessage(t,got,want) 
	})
	t.Run(" empty string defaults to 'world' ",func(t *testing.T){
		got := hello("")
		want := "Hello, World"
		correctMessage(t,got,want)
	})
	
}

func correctMessage(t testing.TB,got,want string){
	t.Helper() // very imp when using reuseable func like this, it points where helper called
	if got != want {
		t.Errorf("got %q want %q",got,want)
	}

}