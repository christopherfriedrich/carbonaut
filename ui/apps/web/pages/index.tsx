import { Button } from '@carbonaut-cloud/ui';
import { ReactElement, useEffect } from 'react';
import { DashboardLayout } from '../layouts/DashboardLayout';
import { LayoutNavigationItem } from '../layouts/types';
import { getResources } from '../integrations/api';

const MainPage = () => {
  useEffect(() => {
    const run = async () => {
      const resources = await getResources();
      const list = resources.getItResourcesList();
      for (const item of list) {
        console.log(item.getLocation());
      }
    };

    run();
  }, []);

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

MainPage.getLayout = (page: ReactElement) => {
  return (
    <DashboardLayout navigation={navigation} title="Dashboard">
      {page}
    </DashboardLayout>
  );
};

export default MainPage;
