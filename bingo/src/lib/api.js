/**
 * @template T
 * @param {T} data
 * @param {string} url
 * @param {string} method
 * @param {unknown} body
 * @returns {Promise<T>}
 */
export const callApi = async (data, url = "", method = "GET", body = null) => {
    const result = await callApiResult(data, url, method, body, true);
    return result.data;
};

/**
 * @template T
 * @param {T} data
 * @param {string} url
 * @param {string} method
 * @param {unknown} body
 * @param {boolean} redirectOnServerError
 * @param {Record<string, string>} additionalHeaders
 * @returns {Promise<{data: T, ok: boolean}>}
 */
export const callApiResult = async (
    data,
    url = "",
    method = "GET",
    body = null,
    redirectOnServerError = false,
    additionalHeaders = {}
) => {
    let token = "";
    try {
        token = localStorage.getItem("token") || "";
    } catch {
        // Authentication storage is best effort in restricted browsers.
    }
    try {
        const response = await fetch(url, {
            method: method,
            headers: {
                "Authorization": `Bearer ${token}`,
                "Content-Type": "application/json",
                ...additionalHeaders,
            },
            body: (body) ? JSON.stringify(body) : null,
        });

        if (response.status === 401) {
            try {
                localStorage.removeItem("token");
            } catch {
                // The failed request still redirects when storage is blocked.
            }
            if (typeof window !== "undefined") {
                window.location.href = "/user/login";
            }
        }

        if (response.status === 404) {
            // window.location.href = "/404";
        }

        if (response.status === 500 && redirectOnServerError) {
            window.location.href = "/500";
        }

        if (response.status === 200) {
            data = await response.json();
        }

        return { data, ok: response.status === 200 };
    } catch (error) {
        console.error("Error:", error);
        return { data, ok: false };
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
    return url;
};