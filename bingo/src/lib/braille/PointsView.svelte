<script>
    import CardPoints from "./CardPoints.svelte";
    import { braille } from "$lib/braille.js";
    import MenuHelperBody from "./MenuHelperBody.svelte";
    import { PointsIcon } from "./icons";
</script>

<MenuHelperBody
    target="pointsview"
    aria_label="Pontos da rodada, do exercício, total de pontos e a série atual"
    title="Pontos"
    icon={PointsIcon(35)}
    color="green"
>
    <div class="container_card">
        <CardPoints
            pointBackgroundColor="blue"
            pointValue={$braille.CurrentClass + 1}
            pointTitle="Série"
            pointTitleHidden="Série"
        />
        <CardPoints
            pointBackgroundColor="orange"
            pointValue={$braille.CurrentRound}
            pointTitle="Rodada"
            pointTitleHidden="Pontos da Rodada"
        />
        <CardPoints
            pointBackgroundColor="green"
            pointValue={$braille.CurrentPunctuation}
            pointTitle="Exercício"
            pointTitleHidden="Pontos do exercício"
        />
        <CardPoints
            pointBackgroundColor="purple"
            pointValue={$braille.TotalPunctuation}
            pointTitle="Total"
            pointTitleHidden="Total de Pontos"
        />
    </div>
</MenuHelperBody>

<style>
    .container_card {
        padding: 0 2rem;
        display: flex;
        justify-content: space-around;
        flex-wrap: wrap;
        gap: 0 4rem;
    }
    @media (max-width: 490px) {
        .container_card {
            margin: 0;
        }
    }
</style>
