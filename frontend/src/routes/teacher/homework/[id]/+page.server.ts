import { api } from '$lib/api';
import type { Actions, PageServerLoad } from './$types';
import type { Solution, Task } from '$lib/types';
import { superValidate } from 'sveltekit-superforms';
import { zod } from 'sveltekit-superforms/adapters';
import { formSchema } from './schema';
import { fail } from '@sveltejs/kit';

export let load: PageServerLoad = async function load({ params, cookies }) {
    const token = cookies.get('teacher_token');
    const res = await api.get(`/solves`, { token });

    const mockSolves: Solution[] = [
        {
            ID: 1,
            CreatedAt: '2024-04-24T21:59:12.537213+07:00',
            UpdatedAt: '2024-04-24T21:59:12.537213+07:00',
            DeletedAt: null,
            Text: 'фывафавы\r\nОднако не все так просто',
            HomeworkID: 1,
            StudentID: 2,
            HomeworkAnswerFiles: null,
            TeacherResumes: null,
            Student: {
                ID: 2,
                CreatedAt: '2024-04-21T23:35:28.053173+07:00',
                UpdatedAt: '2024-04-21T23:35:28.053173+07:00',
                DeletedAt: null,
                Name: 'Артём Майдуров',
                Stage: 11,
                Login: 'artem',
                Lessons: null,
                HomeworksAnswers: null,
                Groups: null
            },
            Homework: {
                ID: 1,
                CreatedAt: '2024-04-21T23:42:51.090281+07:00',
                UpdatedAt: '2024-04-21T23:42:51.090281+07:00',
                DeletedAt: null,
                Description: 'Это тестовый урок',
                LessonID: 1,
                Deadline: '2024-04-18T07:00:00+07:00',
                MaxScore: 10,
                HomeworkFiles: null,
                HomeworkAnswers: null,
                Lesson: null,
            },
        }
    ];

    const mockTask: Task = {
        ID: 2,
        CreatedAt: '2024-04-21T23:45:12.890279+07:00',
        UpdatedAt: '2024-04-21T23:45:12.890279+07:00',
        DeletedAt: null,
        Description: 'Это тестовый урок',
        LessonID: 4,
        Deadline: '2024-04-30T07:00:00+07:00',
        MaxScore: 10,
        HomeworkFiles: null,
        HomeworkAnswers: null,
        Lesson: null,
        IsDone: false,
        GroupId: 2,
        GroupTitle: 'Группа 2: Шизы дрожащие',
    };

    return {
        task: mockTask,
        solutions: mockSolves,
        id: params.id,
        form: await superValidate(zod(formSchema)),
    };
}

export const actions: Actions = {
    default: async (event) => {
        const form = await superValidate(event, zod(formSchema));
        if (!form.valid) {
            return fail(400, { form });
        }

        console.log(form);

        return {
            form,
        };
    },
};


