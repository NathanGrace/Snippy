import PocketBase from 'pocketbase';

export const handle = async ({ event, resolve }) => {
	event.locals.pb = new PocketBase('https://snippy-testing.pockethost.io/')
	event.locals.pb.authStore.loadFromCookie(event.request.headers.get('cookie') || '');

	try {
		if (event.locals.pb.authStore.isValid) {
			await event.locals.pb.collection('users').authRefresh();
			event.locals.user = structuredClone(event.locals.pb.authStore.model);//turns record into POJO
			// event.locals.user!.isAdmin = event.locals.pb.authStore.isAdmin;
		}
	} catch {
		event.locals.pb.authStore.clear();
		event.locals.user = null;
	}

	const response = await resolve(event);

	response.headers.append('set-cookie', event.locals.pb.authStore.exportToCookie());

	return response;
}