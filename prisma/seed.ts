import { PrismaClient } from '@prisma/client';

const prisma = new PrismaClient();

async function main() {
  const vinicius = await prisma.user.upsert({
    where: { email: 'vpcmps@gmail.com' },
    update: {},
    create: {
      name: 'Vinicius',
      email: 'vpcmps@gmail.com',
      password: '123456',
    },
  });
}
main()
  .then(async () => {
    await prisma.$disconnect();
  })
  .catch(async (e) => {
    console.error(e);
    await prisma.$disconnect();
    process.exit(1);
  });
