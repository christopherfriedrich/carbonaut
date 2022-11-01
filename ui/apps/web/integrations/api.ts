import * as grpcWeb from 'grpc-web';
import {
  EmissionDataClient,
  ListITResourcesForProjectRequest,
} from '@carbonaut-cloud/api';

const API_URL = process.env.NEXT_PUBLIC_API_URL!;

if (!API_URL) {
  throw new Error('Missing environment variable `API_URL`.');
}

const service = new EmissionDataClient(API_URL, {});

const request = new ListITResourcesForProjectRequest();
request.setProjectId('123');

export const getResources = async () => {
  const resources = await service.listITResourcesForProject(request, {});
  return resources;
};
