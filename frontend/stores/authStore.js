import { defineStore } from 'pinia';

export const useAuthStore = defineStore('auth', {
    state: () => ({
        user: null,
    }),
    actions: {
        async validate() {
            try {
                const response = await fetch('https://api.alumni-portal.ru/auth/user', {
                    credentials: 'include',
                })
                if (response.ok) {
                    const data = await response.json();
                    this.user = data;
                    return data;
                } else {
                    const errorData = await response.json();
                    throw new Error(errorData.message || 'Login failed');
                }
            } catch (error) {
                console.error('Login error', error);
                throw error;
            }
        },
        async login(email, password) {
            try {
                const response = await fetch('https://api.alumni-portal.ru/auth/login', {
                    method: 'POST',
                    credentials: 'include',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({ email, password }),
                });

                if (response.ok) {
                    const data = await response.json();
                    this.user = data;
                    return data;
                } else {
                    const errorData = await response.json();
                    throw new Error(errorData.message || 'Login failed');
                }
            } catch (error) {
                console.error('Login error', error);
                throw error;
            }
        },
        async logout() {
            try {
                const response = await fetch('https://api.alumni-portal.ru/auth/logout', {
                    method: 'POST',
                    credentials: 'include',
                });

                if (response.ok) {
                    this.user = null;
                } else {
                    const errorData = await response.json();
                    throw new Error(errorData.message || 'Logout failed');
                }
            } catch (error) {
                console.error('Logout error', error);
                throw error;
            }
        },
    },
});
