import { error, type Actions, fail, redirect } from '@sveltejs/kit';
import type { PageServerLoad } from "./$types";
import { api } from '$lib/api';
import type { Lesson, Student } from '$lib/types';

/** @type {PageServerLoad} */
export async function load({ params, cookies }: Parameters<PageServerLoad>[0]) {
    const token = cookies.get('teacher_token');

    const studentsRes = await api.get<{ students: Student[] }>(`/groups/${params.id}/students`, { token });
    const lessonsRes = await api.get<{ students: Lesson[] }>(`/groups/${params.id}/lessons`, { token });

    console.log(studentsRes);

    if (studentsRes.type === 'error' && studentsRes.status === 401 || lessonsRes.type === 'error' && lessonsRes.status === 401) {
        return redirect(303, '/login');
    }

    if (studentsRes.type === "networkerror" || studentsRes.type === "error" || lessonsRes.type === "networkerror" || lessonsRes.type === "error") {
        return redirect(303, '/login');
    }

    const students = studentsRes.data.students as Student[];
    const lessons = lessonsRes.data.students as Lesson[];
    // const students: Student[] = [];

    return {
        students,
        lessons
        // form: await superValidate(zod(formSchema)),
    };
}


// export const actions: Actions = {
//     default: async (event) => {
//         const form = await superValidate(event, zod(formSchema));
//         if (!form.valid) {
//             return fail(400, { form });
//         }

//         const group_id = event.params.id;
//         // const text = form.data.answer;
//         // if (text == undefined || homework_id === undefined) {
//         //     return fail(400, {
//         //         form,
//         //     });
//         // }

//         const token = event.cookies.get('token');
//         const res = await api.get<Student>(`/groups/${group_id}/students`, { token });

//         console.log(res);

//         if (res.type === 'error' && res.status === 401) {
//             return redirect(303, '/login');
//         }

//         if (res.type === "networkerror" || res.type === "error") {
//             return redirect(303, '/login');
//         }

//         // const files = form.data.files.filter((file) => !!file);
//         // if (form.data.files.length > 0) {
//         //     for (const file of files) {
//         //         const fileRes = await fetch(api.url + `/solutions/${res.data.ID}/files`, {
//         //             method: 'POST',
//         //             headers: {
//         //                 'Authorization': `Bearer ${token}`,
//         //                 'Content-Type': file.type,
//         //                 'Content-Disposition': `attachment; filename=${file.name}`,
//         //             },
//         //             body: file,
//         //         });
//         //     }
//         // }

//         // form.data.files = [];

//         return {
//             res,
//         };
//     },
// };


