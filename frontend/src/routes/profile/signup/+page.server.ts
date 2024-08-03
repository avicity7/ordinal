import { PUBLIC_BACKEND_URL } from '$env/static/public'
export const actions = {
  signUp: async ({ request, cookies, locals, fetch }) => {
    const values = await request.formData()
    const Name = values.get("name")
    const Email = values.get("email")
    const Password = values.get("password")
    const captchaRes = values.get("g-recaptcha-response")

    if (Name == "" || Email == "" || captchaRes == "") return { success: false }

    const result = await fetch(`${PUBLIC_BACKEND_URL}:3000/auth/create-user`, {
      method: 'POST',
      body: JSON.stringify({
        Name,
        Email,
        Password,
        Role: 0
      }), 
      headers: {
        'content-type': 'application/json'
      },
    })

    try {
      const parsed = await result.json() as { User: { Name: string, Email: string, RoleID: number, RoleName: string}, AccessToken: string }
      const { User, AccessToken } = parsed
      cookies.set('accessToken', AccessToken, {path: '/'})
      locals.session = User
      return { success: true }
    } catch(err) {
      console.log(err)
    }
  }
}