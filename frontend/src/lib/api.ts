type success<T> = {
    type: "success",
    status: 200,
    data: T
}

type requestError = {
    type: "error",
    status: number,
    error: any
}

type networkError = {
    type: "networkerror"
    error: any
    response?: Response,
}

type ApiResponse<T> = success<T> | networkError | requestError;

interface RequestParams extends RequestInit {
    token?: string;
}

async function apiFetch<T>(url: string, method: string = 'GET', data?: any, options?: RequestParams): Promise<ApiResponse<T>> {
    options = options || {};
    const opts = {
        'Content-Type': 'application/json',
        headers: {},
        ...options,
        method,
    };

    if (opts.token !== undefined) {
        opts.headers = { ...opts.headers, 'Authorization': `Bearer ${opts.token}` };
        delete opts.token;
    }

    if (data !== undefined) {
        opts.body = JSON.stringify(data);
    }

    let res: Response;
    try {
        res = await fetch(url, opts);
    } catch (error) {
        return { type: "networkerror", error };
    }

    if (res.status !== 200) {
        const error = await res.text();
        return { type: "error", status: res.status, error };
    }

    return { type: "success", status: 200, data: await res.json() }
}

const api = {
    url: 'http://localhost:8080',
    post<T>(url: string, data: any, options?: RequestParams) { return apiFetch<T>(this.url + url, 'POST', data, options) },
    get<T>(url: string, options?: RequestParams) { return apiFetch<T>(this.url + url, 'GET', undefined, options) },
}

export { apiFetch, api };

