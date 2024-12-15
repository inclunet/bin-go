<script>
    import PageTitle from "$lib/PageTitle.svelte";
</script>

<PageTitle title="Termos de Uso" />

<section class="container">
    <h2 class="mt-5 mb-4">Termos de Uso do <strong>IncluBingo</strong></h2>
    <section>
        <p>
            Bem-vindo ao IncluBingo! Antes de começar a usar nossos serviços,
            leia atentamente os seguintes Termos de Uso. Ao acessar ou utilizar
            o site IncluBingo, você concorda em cumprir e estar vinculado a
            estes termos. Se você não concorda com esses termos, por favor, não
            utilize o nosso site.
        </p>
    </section>

    <section class="my-4">
        <div class="col-sm-12">
            <div class="card">
                <div class="card-body">
                    <h3 class="card-title">Uso do Site:</h3>
                    <ol>
                        <li>
                            <div class="card-text">
                                <p>
                                    O IncluBingo é um serviço online gratuito,
                                    construído com o propósito de oferecer
                                    oportunidade para pessoas com deficiência se
                                    divertirem em conjunto com outras pessoas
                                    sem deficiência, em conformidade com a
                                    legislação vigente. Fica estritamente
                                    proibida qualquer forma de exploração
                                    comercial deste site, conforme estabelecido
                                    pela Lei 9.981/2000 no Brasil.
                                </p>
                            </div>
                        </li>
                    </ol>
                </div>
            </div>
        </div>

        <div class="col-sm-12">
            <div class="card">
                <div class="card-body">
                    <h3 class="card-title">Responsabilidade e Limitações:</h3>
                    <ol start="2">
                        <li>
                            <div class="card-text">
                                <p>
                                    O IncluBingo não se responsabiliza por
                                    quaisquer danos, perdas ou lesões
                                    resultantes do uso deste site ou da
                                    participação em jogos de bingo recreativos
                                    organizados por você.
                                </p>
                            </div>
                        </li>
                    </ol>
                </div>
            </div>
        </div>

        <div class="col-sm-12">
            <div class="card">
                <div class="card-body">
                    <h3 class="card-title">Alterações nos Termos:</h3>
                    <ol start="1">
                        <li>
                            <div class="card-text">
                                <p>
                                    Reservamo-nos o direito de modificar ou
                                    atualizar estes Termos de Uso a qualquer
                                    momento. Quaisquer alterações serão efetivas
                                    imediatamente após a publicação das versões
                                    revisadas neste site.
                                </p>
                            </div>
                        </li>
                    </ol>
                </div>
            </div>
        </div>

        <div class="col-sm-12">
            <div class="card">
                <div class="card-body">
                    <h3 class="card-title">Lei Aplicável:</h3>
                    <ol start="1">
                        <li>
                            <div class="card-text">
                                <p>
                                    Estes Termos de Uso serão regidos pelas leis
                                    locais do seu país de residência, sem levar
                                    em consideração os conflitos de disposições
                                    legais.
                                </p>
                            </div>
                        </li>
                    </ol>
                </div>
            </div>
        </div>
    </section>

    <section>
        <p>
            Ao continuar a usar o IncluBingo, você concorda com estes Termos de
            Uso. Se você não concorda com eles, por favor, pare de usar o site.
            Obrigado por escolher o IncluBingo para sua diversão inclusiva!
        </p>
    </section>

    <section class="my-4">
        <p>Última atualização: 21 de setembro de 2023</p>
    </section>
</section>

<style>
    :root {
        font-size: 62.5%;
    }
    .card {
        border: none;
    }
    .card-title {
        margin-bottom: 1rem;
        line-height: 4rem;
    }
    .card-text {
        text-align: justify;
    }
    ol {
        padding: 0;
    }
    h3 {
        font-size: 2.2rem;
        margin: 0;
    }
    p,
    section p {
        margin: 0;
        font-size: 1.8rem;
    }
    @media (max-width: 573px) {
        .container {
            padding: 0 2rem 0 2rem;
        }
        h2 {
            line-height: 4rem;
            text-align: center;
        }
        .card {
            border-radius: 1rem;
            box-shadow: 0.5rem 0.5rem 0.5rem #e2e2e2;
        }
        h3 {
            line-height: 4rem;
        }
        p {
            text-align: justify;
            line-height: 3.7rem;
        }
    }
</style>
