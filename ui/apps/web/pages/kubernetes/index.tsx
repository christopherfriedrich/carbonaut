import { Button } from '@carbonaut-cloud/ui';
import { ReactElement } from 'react';
import { DashboardLayout } from '../../layouts/DashboardLayout';
import { LayoutNavigationItem } from '../../layouts/types';
// import { getResources } from '../integrations/api';

const KubernetesPage = () => {
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

KubernetesPage.getLayout = (page: ReactElement) => {
  return (
    <DashboardLayout navigation={navigation} title="Kubernetes">
      {page}
    </DashboardLayout>
  );
};

export default KubernetesPage;
