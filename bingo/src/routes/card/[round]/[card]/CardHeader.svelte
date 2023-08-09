<script>
    import Autoplay from "./Autoplay.svelte";
import Card from "./Card.svelte";
    import NewRound from "./NewRound.svelte";

export let card = {Card: 0, Checked: 0, LastNumber: 0, Round: 0};
</script>

<div>
    <div>
        <p>{card.Checked} Bolas Sorteadas</p>
    </div>
    <div>
<Autoplay {card} />
<NewRound {card} />
    </div>

    <div aria-live="polite">
        <p>Última bola sorteada: {card.LastNumber}</p>
    </div>
</div>
