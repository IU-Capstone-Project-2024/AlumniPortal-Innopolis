import { defineStore } from 'pinia'
export const useOfferStore = defineStore({
    id: 'offer-store',
    state: () => {
        return {
            offers: [
                {
                    id: 1,
                    url: 'iu-building.png',
                    title: 'Inno',
                    desc: 'description',
                    collected: 50,
                    goal: 150,
                },
                {
                    id: 2,
                    url: 'iu-building.png',
                    title: 'Inno',
                    desc: 'description',
                    collected: 50,
                    goal: 150,
                },
                {
                    id: 3,
                    url: 'iu-building.png',
                    title: 'Inno',
                    desc: 'description',
                    collected: 50,
                    goal: 150,
                },
                {
                    id: 4,
                    url: 'iu-building.png',
                    title: 'Inno',
                    desc: 'description',
                    collected: 50,
                    goal: 150,
                },
            ]
        }
    },
    actions: {
        addOffer(offer) {
            this.offers.push(offer)
        }
    },
    getters: {
        getOffers: state => state.offers,
        getOfferById: (state) => {
            return (id) => state.offers.find(o => o.id === id)},
    }
})