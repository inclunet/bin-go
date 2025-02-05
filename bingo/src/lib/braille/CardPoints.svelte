<script>
    export let pointBackgroundColor;
    export let pointValue;
    export let pointTitle;

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
                nameColor = "#8A2BE2";
            default:
                break;
        }
        return nameColor;
    };
</script>

<div class="card text-center" style="width: 18rem;">
    <div class="card-body">
        <div class="card-points" style="background-color: {choiceColor()};">
            <p>{pointValue}</p>
        </div>
        <h4 class="card-title">{pointTitle}</h4>
    </div>
</div>

<style>
    :root {
        --white: #fff;

        font-size: 62.5%;
    }
    h4 {
        font-size: 2rem;
    }

    .card {
        border: 1px solid rgb(102, 102, 102);
    }
    .card-body {
        padding: 0;
    }

    .card-points {
        font-weight: bold;
    }
    .card-points p {
        font-size: 4rem;
        color: var(--white);
        padding: 25px 0;
    }

    .card-title {
        font-weight: 600;
        padding: 8px;
    }
</style>
