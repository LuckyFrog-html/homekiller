type success<T> = {
    type: "success",
    status: 200,
    data: T
}

type requestError = {
    type: "error",
    status: number
}

type networkError = {
    type: "networkerror"
    error: any
    response?: Response,
}

type ApiResponse<T> = success<T> | networkError | requestError;


async function apiFetch<T>(url: string, method: string = 'GET', data?: any, options?: RequestInit): Promise<ApiResponse<T>> {
    options = options || {};
    const opts = {
        ...options,
        method,
        'Content-Type': 'application/json',
    };
    if (data !== undefined) {
        opts.body = JSON.stringify(data);
    }
    let res: Response;
    try {
        res = await fetch(url, opts);
    } catch (e) {
        return { type: "networkerror", error: e };
    }

    if (!res.ok) {
        return { type: "networkerror", error: "idk", response: res };
    }

    if (res.status !== 200) {
        return { type: "error", status: res.status };
    }

    return { type: "success", status: 200, data: await res.json() }
}

const api = {
    url: 'http://localhost:8080',
    post<T>(url: string, data: any, options?: RequestInit) { return apiFetch<T>(this.url + url, 'POST', data, options) },
    get<T>(url: string, options?: RequestInit) { return apiFetch<T>(this.url + url, 'GET', undefined, options) },
}

export { apiFetch, api };
