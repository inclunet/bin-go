<script>
    import PageTitle from "$lib/PageTitle.svelte";
</script>

<PageTitle title="Configuração do NVDA" game="Inclubraille" />

<div class="container">
    <h2>Como configurar o NVDA para Jogar Inclubraille</h2>

    <section>
        <p>
            Antes de iniciar o jogo Inclubraille, é necessário ajustar a
            configuração de verbosidade do NVDA para garantir que os caracteres
            Braille sejam corretamente anunciados.
        </p>
    </section>

    <section>
        <p>Passos:</p>

        <ol class="list-group list-group-numbered">
            <li class="list-group-item">
                Abra o NVDA e pressione NVDA + N para acessar o menu do leitor
                de telas.
            </li>
            <li class="list-group-item">
                Navegue até Preferências e entre em Configurações.
            </li>
            <li class="list-group-item">
                No painel de configurações, vá até a categoria Voz.
            </li>
            <li class="list-group-item">
                Encontre a opção Verbosidade dos caracteres e símbolos.
            </li>
            <li class="list-group-item">Selecione a opção Toda.</li>
            <li class="list-group-item">
                Pressione OK para salvar as configurações.
            </li>
        </ol>
    </section>

    <section class="mt-5">
        <p>Agora o NVDA lerá corretamente os caracteres Braille do jogo.</p>

        <p>Boa diversão!</p>
    </section>
</div>

<style>
    li {
        font-size: 1.8rem;
    }
</style>
