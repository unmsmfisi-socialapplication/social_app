import axios from 'axios'
import MockAdapter from 'axios-mock-adapter'
import { doGet } from './../../../data/api/apiService' // Asegúrate de ajustar la ruta según la estructura de tu proyecto

const mock = new MockAdapter(axios)

describe('API Requests', () => {
    const url = 'https://rickandmortyapi.com/api'

    afterEach(() => {
        mock.reset()
    })

    it('should make a GET request successfully', async () => {
        const responseData = { data: 'Mocked data' }
        mock.onGet(url).reply(200, responseData)

        const response = await doGet(url)
        expect(response.status).toBe(200)
        expect(response.data).toEqual(responseData)
    })

    it('should handle errors', async () => {
        mock.onGet(url).reply(404)

        await expect(doGet(url)).rejects.toThrowError()
    })
})
