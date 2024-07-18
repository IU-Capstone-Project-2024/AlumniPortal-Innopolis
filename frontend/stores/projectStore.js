import { defineStore } from 'pinia'
export const useProjectStore = defineStore({
    id: 'project-store',
    state: () => {
        return {
            projects: [
                {
                    id: 1,
                    url: 'inno.png',
                    title: 'Inno',
                    desc: 'description',
                    collected: 50,
                    goal: 150,
                },
                {
                    id: 2,
                    url: 'inno.png',
                    title: 'Inno',
                    desc: 'description',
                    collected: 50,
                    goal: 150,
                },
                {
                    id: 3,
                    url: 'inno.png',
                    title: 'Inno',
                    desc: 'description',
                    collected: 50,
                    goal: 150,
                },
                {
                    id: 4,
                    url: 'inno.png',
                    title: 'Inno',
                    desc: 'description',
                    collected: 50,
                    goal: 150,
                },
                {
                    id: 5,
                    url: 'inno.png',
                    title: 'Inno',
                    desc: 'description',
                    collected: 50,
                    goal: 150,
                },
                {
                    id: 6,
                    url: 'inno.png',
                    title: 'Inno',
                    desc: 'description',
                    collected: 50,
                    goal: 150,
                },
                {
                    id: 7,
                    url: 'inno.png',
                    title: 'Inno',
                    desc: 'description',
                    collected: 50,
                    goal: 150,
                },
                {
                    id: 8,
                    url: 'inno.png',
                    title: 'Inno',
                    desc: 'description',
                    collected: 50,
                    goal: 150,
                },
                {
                    id: 9,
                    url: 'inno.png',
                    title: 'Inno',
                    desc: 'description',
                    collected: 50,
                    goal: 150,
                },
                {
                    id: 10,
                    url: 'inno.png',
                    title: 'Inno',
                    desc: 'description',
                    collected: 50,
                    goal: 150,
                },
            ]
        }
    },
    actions: {
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