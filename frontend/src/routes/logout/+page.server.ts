import { redirect, type Actions } from "@sveltejs/kit";

export const actions: Actions = {
    default: async ({ cookies }) => {
        cookies.delete('token', { path: '/', secure: false });
        cookies.delete('teacher_token', { path: '/', secure: false });
        return redirect(303, '/login');
    }
}
