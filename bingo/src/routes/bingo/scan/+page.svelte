<script>
    import { onMount } from "svelte";

    /**
     * @type {HTMLVideoElement}
     */
    let videoElement;
    /**
     * @type {HTMLCanvasElement}
     */
    let canvasElement;
    let photoTaken = false;
    /**
     * @type {MediaStream | null}
     */
    let stream;
    let facingMode = "environment";

    const startCamera = async () => {
        try {
            stream = await navigator.mediaDevices.getUserMedia({
                video: {
                    facingMode: facingMode,
                },
            });

            videoElement.srcObject = stream;
                    } catch (error) {
            console.error(error);
        }
    };

    const toggleCamera = async () => {
        if (facingMode === "environment") {
            facingMode = "user";
        } else {
            facingMode = "environment";
        }

        if (stream != null) {
            stream
                .getTracks()
                .forEach((/** @type {{ stop: () => void; }} */ track) => {
                    track.stop();
                });
        }

        startCamera();
    };

    const takePhoto = () => {
        const context = canvasElement.getContext("2d");

        if (context != null) {
            context.drawImage(videoElement, 0, 0, 640, 480);
            photoTaken = true;
        }
    };

    const uploadPhoto = async () => {
        const dataUrl = canvasElement.toDataURL("image/jpeg");
        const response = await fetch("/api/upload", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({ image: dataUrl }),
        });
        const result = await response.json();
        console.log("Resposta da API:", result);
    };

    onMount(startCamera);
</script>
<h2>Digitalise sua cartela física</h2>
<video bind:this={videoElement} autoplay playsinline aria-hidden="true"></video>
<canvas bind:this={canvasElement} style="display: none;"></canvas>

{#if !photoTaken}
    <button on:click={takePhoto}>Tirar Foto</button>
    <button on:click={toggleCamera}>Alternar Câmera</button>
{/if}

{#if photoTaken}
    <button on:click={uploadPhoto}>Enviar Foto</button>
    <button on:click={() => (photoTaken = false)}>Tirar Novamente</button>
{/if}
