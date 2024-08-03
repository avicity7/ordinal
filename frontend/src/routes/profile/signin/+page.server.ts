import { PUBLIC_BACKEND_URL } from '$env/static/public';

export const actions = {
	login: async ({ request, cookies, locals }) => {
		const values = await request.formData();
		const email = values.get('email');
		const password = values.get('password');
		const result = await fetch(`${PUBLIC_BACKEND_URL}:3000/auth/login`, {
			method: 'POST',
			body: JSON.stringify({
				Email: email,
				Password: password
			}),
			headers: {
				'content-type': 'application/json'
			}
		});

		try {
			const parsed = (await result.json());
			const { User, AccessToken } = parsed;
			cookies.set('accessToken', AccessToken, { path: '/' });
			locals.session = User;
			return {
				success: true,
				account: User
			};
		} catch {
			return { success: false }
		}
	},
	signOut: async ({ cookies, locals }) => {
		cookies.delete('accessToken', { path: '/' });
		locals.session = null;
	}
};