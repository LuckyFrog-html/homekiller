import { api } from '$lib/api';
import type { RequestHandler } from '@sveltejs/kit';

export let GET: RequestHandler = async function GET({ url }) {
    let res: Response;
    try {
        res = await fetch(api.url + url.pathname);
    } catch (error) {
        return new Response("Not found", { status: 404 });
    }

    if (res.status !== 200) {
        const error = await res.text();
        return new Response(error, { status: res.status });
    }

    return res;
}
