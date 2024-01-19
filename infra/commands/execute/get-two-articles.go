package CommandExecute

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/bwmarrin/discordgo"
)

type ResponseData struct {
	Value []Article `json:"value"`
	Count int       `json:"count"`
}

type Response struct {
	Message string       `json:"message"`
	Data    ResponseData `json:"data"`
	Status  int          `json:"status"`
}

type Article struct {
	ArtigoID             int      `json:"artigoID"`
	ArtigoUUID           string   `json:"artigoUUID"`
	ArtigoAutorID        int      `json:"artigoAutorID"`
	ArtigoTitle          string   `json:"artigoTitle"`
	ArtigoSubtitle       string   `json:"artigoSubtitle"`
	ArtigoSubdescription string   `json:"artigoSubdescription"`
	ArtigoBody           string   `json:"artigoBody"`
	ArtigoType           string   `json:"artigoType"`
	Publicado            bool     `json:"publicado"`
	CreatedAt            string   `json:"createdAt"`
	ArtigoStars          []string `json:"ArtigoStars"`
	ArtigoImages         []struct {
		ArtigoImagesID  int    `json:"artigoImagesID"`
		ArtigoUUID      string `json:"artigoUUID"`
		ArtigoImageURL  string `json:"artigoImageURL"`
		ArtigoImageName string `json:"artigoImageName"`
		CreatedAt       string `json:"createdAt"`
	} `json:"ArtigoImages"`
	Usuario struct {
		UsuarioNome         string   `json:"usuarioNome"`
		UsuarioEmail        string   `json:"usuarioEmail"`
		UsuarioContaAtivada bool     `json:"usuarioContaAtivada"`
		UsuarioPermissions  []string `json:"usuarioPermissions"`
		CreatedAt           string   `json:"createdAt"`
		UpdatedAt           string   `json:"updatedAt"`
		UsuarioNascimento   string   `json:"usuarioNascimento"`
		UsuarioTelefone     string   `json:"usuarioTelefone"`
	} `json:"Usuario"`
}

func GetTwoArticles(session *discordgo.Session, interation *discordgo.InteractionCreate) {
	url := "https://api.seliganamidia.xyz/articles?pageIndex=1&pageSize=2"

	response, err := http.Get(url)

	if err != nil {
		log.Println("Error: ", err)

		session.InteractionRespond(interation.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Erro ao buscar os artigos",
			},
		})
		return
	}
	defer response.Body.Close()

	body, erro := ioutil.ReadAll(response.Body)
	if erro != nil {
		fmt.Println("Erro ao ler o corpo da resposta:", erro)

		session.InteractionRespond(interation.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Erro ao buscar os artigos",
			},
		})

		return
	}

	var articleStruct Response

	erro = json.Unmarshal(body, &articleStruct)
	if erro != nil {
		fmt.Println("Erro ao decodificar JSON:", erro)
		return
	}

	log.Println(string(articleStruct.Message))
	log.Println(articleStruct.Status)
	log.Println(articleStruct.Data.Count)

	if articleStruct.Data.Count == 0 {
		session.InteractionRespond(interation.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Não há artigos",
			},
		})
		return
	}

	first := articleStruct.Data.Value[0]
	second := articleStruct.Data.Value[1]

	phrase := "Titulo: " + first.ArtigoTitle + "\n" + "Subtitulo: " + first.ArtigoSubtitle + "\n" + "Link: https://seliganamidia.xyz/jornal/" + first.ArtigoUUID + "\n --------------------------------------------------- \n"
	phrase2 := "Titulo: " + second.ArtigoTitle + "\n" + "Subtitulo: " + second.ArtigoSubtitle + "\n" + "Link: https://seliganamidia.xyz/jornal/" + second.ArtigoUUID + "\n\n"

	session.InteractionRespond(interation.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: string(phrase + phrase2),
		},
	})

}
