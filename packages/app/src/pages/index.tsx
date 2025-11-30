import { Divider } from '@status/components/Divider';
import { Linear } from '@status/components/Linear';
import { Navbar } from '@status/components/Navbar';
import { Status } from '@status/components/Status';

const StatusPage = () => {
  return (
    <div className="min-h-screen">
      <Linear.Background />
      <div className="relative z-10 flex flex-col">
        <Navbar />
        <Divider />
        <div className="container mx-auto grow overflow-auto p-4 md:p-8">
          <Status />
        </div>
      </div>
    </div>
  );
};

export default StatusPage;
