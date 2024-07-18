import { defineStore } from "pinia";
import settings from '../settings.json'
const path = settings.url

export const useFullStore = defineStore('FullStore', {
    state: () => ({
        fullTicket: [],
        comments: [],
        count: 0,
    }),
    actions:{
        async getFullTicket(token, id) {
            const res = await fetch(path+'tsr/fulltsr',{
                method: "POST",
                body: JSON.stringify({
                    token: token,
                    tsr_id: id,
                }),
            })
            const data = await res.json()
            if (res.status === 200) {
                console.log(data)
                this.fullTicket = data
            }
        },
        async getTicketComments(token, id) {
            const res = await fetch(path+'tsr/comments',{
                    method: "POST",
                    body: JSON.stringify({
                        token: token,
                        tsr_id: id,
                    }),
                })
                const data = await res.json()
                if (res.status === 200) {
                    console.log(data)
                    this.comments = data.comments
                    this.count = data.count
            }
        },
        async sendTicketComment(token, id, text) {
            const res = await fetch(path+'tsr/comment',{
                    method: "POST",
                    body: JSON.stringify({
                        token: token,
                        tsr_id: id,
                        comment_text: text,
                    }),
                })
                if (res.status === 200) {
                    this.getTicketComments(token, id)
            }
        },
        async setEmployee(tsrid, emplid, token) {
            const res = await fetch(path+'tsr/employee',{
                method: "POST",
                body: JSON.stringify({
                    tsr_uuid: tsrid,
                    employee_uuid: emplid,
                    token: token
                })
            })
            if (res.status === 200) {
                this.getFullTicket(token, tsrid)
            }
        },
        async finishTstr(tsrid, token) {
            const res = await fetch(path+'tsr/finish',{
                method: "POST",
                body: JSON.stringify({
                    token: token,
                    tsr_uuid: tsrid,
                })
            })
            if (res.status === 200) {
                this.getFullTicket(token, tsrid)
            }
        },
        async toggleImportance(tsrid, token, importance) {
            const res = await fetch(path+'tsr/importance',{
                method: "POST",
                body: JSON.stringify({
                    token: token,
                    tsr_uuid: tsrid,
                    important: importance,
                })
            })
            if (res.status === 200) {
                this.getFullTicket(token, tsrid)
            }
        },
        async applyTstr(tsrid, token) {
            const res = await fetch(path+'tsr/apply',{
                method: "POST",
                body: JSON.stringify({
                    token: token,
                    tsr_uuid: tsrid,
                })
            })
            if (res.status === 200) {
                return true
            }
        }
    }, 
})