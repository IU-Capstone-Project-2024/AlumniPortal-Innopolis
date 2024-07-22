import { defineStore } from 'pinia'
export const useProjectStore = defineStore({
    id: 'project-store',
    state: () => {
        return {
            projects: []
        }
    },
    actions: {
        async fetchProjects() {
            this.projects = []
            try {
                const response = await fetch('https://api.alumni-portal.ru/projects', {
                    credentials: "include",
                })
                if (response.ok) {
                    const data = await response.json();
                    for (const project of data) {
                        this.addProject({
                            id: project.ID,
                            url: 'inno.png',
                            title: project.Name,
                            desc: project.Description,
                            collected: project.Collected,
                            goal: project.Goal,
                            Status: project.Status,
                        })
                    }
                } else {
                    const errorData = await response.json();
                    throw new Error(errorData.message || 'Project fetch failed');
                }
            } catch (error) {
                console.error('Fetch prj error', error);
                throw error;
            }
        },
        addProject(project) {
            this.projects.push(project)
        }
    },
    getters: {
        getProjects: state => state.projects,
        getVerifiedProjects: state => {
            return state.projects.filter(p => p.Status === 'Accepted')
        },
        getProjectById: (state) => {
            return (id) => state.projects.findOne(p => p.id === id)},
    }
})