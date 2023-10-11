import {Service} from "@/services";

export function imageUpdateApi(data) {
    return Service.put('/api/images/', data)
}

export const imageRemoveBatchApi = (id_list) => {
    return Service.delete('/api/images/', {data: {id_list}})
}