// import { error, redirect } from '@sveltejs/kit';
//
// export const actions = {
// 	register: async ({ locals, request }) => {
// 		const body = Object.fromEntries(await request.formData());//I need to setup form data
//
// 		try {
// 			await locals.pb.collection('users').create({ ...body });
// 			// await locals.pb.collection('users').requestVerification(body.email);
// 		} catch (err) {
// 			console.log('Error: ', err);
// 			throw error(500, 'Something went wrong');
// 		}
//
// 		throw redirect(303, '/login');
// 	}
// };


import { redirect } from '@sveltejs/kit';

export const actions = {
	login: async ({ request, locals }) => {
		const formData = await request.formData();
		const username = formData.get('username') as string;
		const password = formData.get('password') as string;
		console.log(username, password);
		try {
			await locals.pb.collection('staff_members').authWithPassword(username, password);

			if (!locals.pb.authStore.isValid) {
				locals.pb.authStore.clear();
			}
			// if (!locals.pb?.authStore?.model?.verified) {
			// 	console.log("not verified");
			// 	locals.pb.authStore.clear();
			// 	return {
			// 		notVerified: true
			// 	};
			// } //else verified successfully

		//TODO improve error handling
		} catch (err) {
			console.log('Error:', err );
		}

		redirect(303, '/calendar');
	}
};