declare global {
	declare namespace App {
		interface Locals {
			session: { Name: string, Email: string, RoleID: number, RoleName: string } | null,
		}
	}
}

export {}