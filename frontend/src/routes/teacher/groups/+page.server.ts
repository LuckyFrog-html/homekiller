import { api } from "$lib/api";
import { redirect } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";
import type { Task } from "$lib/types";

// /** @type {PageServerLoad} */
// export async function load({ cookies }: Parameters<PageServerLoad>[0]): Promise<{ tasks: Task[] }> {
//     const token = cookies.get('token');
//     const req = await api.get<any>('/teacher/groups', { token });

//     console.log(req);

//     if (req.type === "success") {
//         // const groups = req.data.groups as Task[] || [];
//         // return { groups };
//     }

//     // return redirect(302, '/login');
// }
