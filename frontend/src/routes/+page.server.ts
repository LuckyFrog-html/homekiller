import { redirect } from '@sveltejs/kit';

export function load({ cookies }) {
    const token = cookies.get('token');
    const teacher_token = cookies.get('teacher_token');

    if (token) {
        return redirect(303, '/student/tasks');
    }

    if (teacher_token) {
        return redirect(303, '/teacher/groups');
    }

    return redirect(303, '/login');
}
