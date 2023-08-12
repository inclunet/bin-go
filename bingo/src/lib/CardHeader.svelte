<script>
    import Autoplay from "./Autoplay.svelte";
    import Draw from "./Draw.svelte";
    import NewRound from "./NewRound.svelte";

    export let card = {
        Autoplay: false,
        Card: 0,
        Checked: 0,
        LastNumber: 0,
        Round: 0,
    };
</script>

<div>
    <div>
        <p>{card.Checked} Bolas Sorteadas</p>
    </div>
    <div>
        <Autoplay bind:card={card} />
        <NewRound bind:card={card} />
        <Draw bind:card />
    </div>

    <div aria-live="polite">
        <p>Última bola sorteada: {card.LastNumber}</p>
    </div>
</div>
