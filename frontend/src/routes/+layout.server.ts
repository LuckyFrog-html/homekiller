import type { LayoutServerLoad } from './$types';

export const load: LayoutServerLoad = async function({ cookies }) {
    const token = cookies.get('token');
    const teacher_token = cookies.get('teacher_token');
    const isLogined = !!token || !!teacher_token;
    return { isLogined };
}
