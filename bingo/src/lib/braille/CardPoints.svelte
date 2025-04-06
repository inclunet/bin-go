<script>
    export let pointBackgroundColor;
    export let pointValue;
    export let pointTitle;
    export let pointTitleHidden = "";

    const choiceColor = () => {
        let nameColor = "";
        switch (pointBackgroundColor) {
            case "blue":
                nameColor = "#1E90FF";
                break;
            case "orange":
                nameColor = "#FF4500";
                break;
            case "green":
                nameColor = "#42CD42";
                break;
            case "purple":
                nameColor = "#954ed6";
            default:
                break;
        }
        return nameColor;
    };
</script>

<div class="card">
    <div class="card-body">
        <div class="card-points" style="background-color: {choiceColor()};">
            <p class="card-title hidden-area">{pointTitleHidden}</p>
            <p>{pointValue}</p>
        </div>
        <h4 aria-hidden="true" class="card-title">{pointTitle}</h4>
    </div>
</div>

<style>
    h4 {
        font-size: 1.8rem;
    }
    .card {
        height: 15rem;
        border: none;
    }

    .card-body {
        display: flex;
        flex-direction: column;
        align-items: center;
    }

    .card-points {
        display: flex;
        justify-content: center;
        align-items: center;
        width: 7.8rem;
        height: 7.8rem;
        border-radius: 8%;
        padding: 0;
    }

    .card-points p {
        display: flex;
        align-items: center;
        margin: 0;
        padding: 0;
        font-weight: 500;
        font-size: 5rem;
    }

    .card-title {
        padding: 1rem 0;
    }
</style>
