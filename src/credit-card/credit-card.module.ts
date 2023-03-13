import { Module } from '@nestjs/common';
import { CreditCardService } from './credit-card.service';
import { CreditCardController } from './credit-card.controller';

@Module({
  imports: [
    MulterModule.register({
      dest: './files',
    }),
  ],
  controllers: [CreditCardController],
  providers: [CreditCardService],
})
export class CreditCardModule {}
