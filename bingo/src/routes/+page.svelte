<script>
    import { goto } from "$app/navigation";
    import { page } from "$app/stores";

    function getEndpointUrl(call = "") {
        if ($page.url.port == "5173") {
            return (
                $page.url.protocol +
                "//" +
                $page.url.hostname +
                ":8080/api" +
                call
            );
        } else {
            return $page.url.protocol + "//" + $page.url.host + "/api" + call;
        }
    }

    async function new75Card() {
        let card = { Round: 0 };
        const url = getEndpointUrl("/card/0/75");
        const response = await fetch(url);
        card = await response.json();
        goto("/card/" + card.Round);
    }
</script>

<h1>Bem-Vindo ao Inclubingo</h1>
<p>Vamos Jogar Escolha a quantidade de bolinhas que serão sorteadas:</p>
<ul>
    <li>
        <button on:click={new75Card} class="btn btn-primary"
            >Bingo com 75 bolinhas</button
        >
    </li>
</ul>

<style>
    h1,
    p {
        text-align: center;
    }
    button {
        display: block;
        margin: 0 auto;
    }
</style>
