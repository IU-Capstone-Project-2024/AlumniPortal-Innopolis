import { defineStore } from 'pinia'
export const useProjectStore = defineStore({
    id: 'project-store',
    state: () => {
        return {
            projects: null
        }
    },
    actions: {
        async fetchProjects() {
            try{
                const res = fetch('https://api.alumni-portal.ru/projects', {
                    credentials: "include",
                })
                if (response.ok) {
                   const data = await res.json();
                   console.log(data);
                }
            } catch (error) {
                console.error('Fetch project error', error);
            }
        },
        addProject(project) {
            this.projects.push(project)
        }
    },
    getters: {
        getProjects: state => state.projects,
        getProjectById: (state) => {
            return (id) => state.projects.find(p => p.id === id)},
    }
})