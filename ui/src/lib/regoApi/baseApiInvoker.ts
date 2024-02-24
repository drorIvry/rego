import { env } from "~/env";
import { AxiosResponse, Axios, AxiosError } from "axios";

const axios = require("axios").default;

export class RegoApi {
  regoBaseUrl: string;
  regoClient: Axios;

  constructor() {
    this.regoBaseUrl = env.REGO_URL;

    if (this.regoBaseUrl.endsWith("/")) {
      this.regoBaseUrl = this.regoBaseUrl.substring(
        0,
        this.regoBaseUrl.length - 1,
      );
    }

    this.regoClient = axios.create({
      baseUrl: this.regoBaseUrl,
    });
  }

  protected async invokeApi(
    method: string,
    endpoint: string,
    api_key: string,
    body?: string,
    parameters?: { [key: string]: any },
    headers?: { [key: string]: string },
  ): Promise<AxiosResponse> {
    endpoint = endpoint.startsWith("/") ? endpoint : "/" + endpoint;
    headers = headers ? headers : {};

    const response: AxiosResponse = await this.regoClient.request({
      method: method.toLowerCase(),
      url: this.regoBaseUrl + endpoint,
      data: body,
      params: parameters,
      headers: {
        "X-Rego-Api-Key": api_key,
        ...headers,
      },
    });

    return response;
  }
}
