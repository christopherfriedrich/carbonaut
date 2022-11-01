import { Button } from '@carbonaut-cloud/ui';
import { ReactElement } from 'react';
import { DashboardLayout } from '../../layouts/DashboardLayout';
import { LayoutNavigationItem } from '../../layouts/types';
// import { getResources } from '../integrations/api';

const ResourcesPage = () => {
  // useEffect(() => {
  //   const run = async () => {
  //     const resources = await getResources();
  //     console.log(resources);
  //   };

  //   run();
  // }, []);

  return (
    <div>
      <h1>Web</h1>
      <Button />
    </div>
  );
};

const navigation: LayoutNavigationItem[] = [
  { name: 'Dashboard', pathname: '/' },
  { name: 'Resources', pathname: '/resources' },
  { name: 'Kubernetes', pathname: '/kubernetes' },
];

ResourcesPage.getLayout = (page: ReactElement) => {
  return (
    <DashboardLayout navigation={navigation} title="Resources">
      {page}
    </DashboardLayout>
  );
};

export default ResourcesPage;
