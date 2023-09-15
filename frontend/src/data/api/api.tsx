import axios, { AxiosResponse } from 'axios';
  export async function doGet<T>(
    url: string,
    params: Record<string, any> = {},
    header: Record<string, any> = {}
  ): Promise<AxiosResponse<T>> {
    const idToken = localStorage.getItem("token") || "";


    try {
      const response = await axios({
        url,
        method: "get",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${idToken}`,
          ...header,
        },
        data: params,
      });
      return response;
    } catch (error) {
      console.error("error", error);
      throw error; //TODO: Re-lanza el error para que pueda ser manejado en otros lugares si es necesario.
    }
  }

  export async function doPost<T>(
    url: string,
    params: Record<string, any> = {}
  ): Promise<AxiosResponse<T>> {
    const idToken = localStorage.getItem('token') || '';
  
    try {
      const response = await axios({
        url,
        method: 'post',
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${idToken}`,
        },
        data: params,
      });
      return response;
    } catch (error) {
      console.error('errorDoPost', error);
      throw error; //TODO: Re-lanza el error para que pueda ser manejado en otros lugares si es necesario.
    }
  }

  export async function doPut<T>(
    url: string,
    params: Record<string, any> = {},
    header: Record<string, any> = {}
  ): Promise<AxiosResponse<T>> {
    const idToken = localStorage.getItem('token') || '';
  
    try {
      const response = await axios({
        url,
        method: 'put',
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${idToken}`,
          ...header,
        },
        data: params,
      });
      return response;
    } catch (error) {
      console.error('errorDoPut', error);
      throw error; //TODO: Re-lanza el error para que pueda ser manejado en otros lugares si es necesario.
    }
  }
  export async function doDelete<T>(
    url: string,
    params: Record<string, any> = {}
  ): Promise<AxiosResponse<T>> {
    const idToken = localStorage.getItem('token') || '';
  
    try {
      const response = await axios({
        url,
        method: 'delete',
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${idToken}`,
        },
        data: params,
      });
      return response;
    } catch (error) {
      console.error('errorDoDelete', error);
      throw error; //TODO: Re-lanza el error para que pueda ser manejado en otros lugares si es necesario.
    }
  }