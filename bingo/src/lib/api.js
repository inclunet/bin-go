import { error } from "@sveltejs/kit";

export const callApi = async (data = {}, url = "", method = "GET", body = null) => {
    const token = (localStorage.getItem("token")) ? localStorage.getItem("token") : "";
    console.log(url);
    console.log(method)
    console.log(body)
    try {
        const response = await fetch(url, {
            method: method,
            headers: {
                "Authorization": `Bearer ${token}`,
                "Content-Type": "application/json",
            },
            body: (body) ? JSON.stringify(body) : null,
        });

        if (response.status === 401) {
            localStorage.removeItem("token");
            window.location.href = "/user/login";
        }

        if (response.status === 404) {
            // window.location.href = "/404";
        }

        if (response.status === 500) {
            window.location.href = "/500";
        }

        if (response.status === 200) {
            data = await response.json();
        }

        return data;
    } catch (error) {
        console.error("Error:", error);
        return data;
    }
};

export const getWSEndpoint = (path = "") => {
    let url = document.location.href;

    if (document.location.protocol === "https:") {
        url = url.replace(document.location.protocol, "wss:");
    } else {
        url = url.replace(document.location.protocol, "ws:");
    }

    url = url.replace(document.location.pathname, path);
    console.log(url);
    return url;
};