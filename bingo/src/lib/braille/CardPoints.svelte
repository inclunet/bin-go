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
                nameColor = "#954ed6";
            default:
                break;
        }
        return nameColor;
    };
</script>

<div class="card text-center">
    <div class="card-body">
        <div class="card-points" style="background-color: {choiceColor()};">
            <p>{pointValue}</p>
        </div>
        <h4 class="card-title">{pointTitle}</h4>
    </div>
</div>

<style>
    :root {
        --black: #000;
        --white: #fff;

        font-size: 62.5%;
    }
    h4 {
        font-size: 2rem;
    }

    .card {
        width: 18rem;
        border: 1px solid rgb(102, 102, 102);
    }
    .card-body {
        padding: 0;
    }

    .card-points {
        display: flex;
        justify-content: center;
        align-items: center;
        height: 10rem;
        font-weight: bold;
    }
    .card-points p {
        margin: 0;
        padding: 0;
        font-size: 4rem;
        color: var(--black);
    }

    .card-title {
        font-weight: 600;
        padding: 8px;
    }

    @media (max-width: 490px) {
        h4 {
            font-size: 1.8rem;
        }
        .card {
            display: flex;
            align-items: center;
            width: 8rem;
            height: 13rem;
            border: none;
            flex-wrap: wrap;
        }

        .card-points {
            display: flex;
            justify-content: center;
            align-items: center;
            width: 7.8rem;
            height: 7.8rem;
            border-radius: 50%;
            padding: 0;
        }

        .card-points p {
            padding: 0;
        }

        .card-title {
            padding: 1rem 0;
        }
    }
</style>
