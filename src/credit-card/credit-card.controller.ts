import {
  Controller,
  Get,
  Post,
  Body,
  Patch,
  Param,
  Delete,
  UseInterceptors,
  UploadedFile,
} from '@nestjs/common';
import { CreditCardService } from './credit-card.service';
import { CreateCreditCardDto } from './dto/create-credit-card.dto';
import { UpdateCreditCardDto } from './dto/update-credit-card.dto';
import { FileInterceptor } from '@nestjs/platform-express';
import * as csv from 'fast-csv';
import { ApiTags } from '@nestjs/swagger';
@ApiTags('credit-card')
@Controller('credit-card')
export class CreditCardController {
  constructor(private readonly creditCardService: CreditCardService) {}

  @Post('upload')
  @UseInterceptors(FileInterceptor('file'))
  uploadFile(@UploadedFile() file: Express.Multer.File) {
    const response = {
      originalname: file.originalname,
      filename: file.filename,
    };
    return response;
  }

  @Post()
  create(@Body() createCreditCardDto: CreateCreditCardDto) {
    return this.creditCardService.create(createCreditCardDto);
  }

  @Get()
  findAll() {
    return this.creditCardService.findAll();
  }

  @Get(':id')
  findOne(@Param('id') id: string) {
    return this.creditCardService.findOne(+id);
  }

  @Patch(':id')
  update(
    @Param('id') id: string,
    @Body() updateCreditCardDto: UpdateCreditCardDto,
  ) {
    return this.creditCardService.update(+id, updateCreditCardDto);
  }

  @Delete(':id')
  remove(@Param('id') id: string) {
    return this.creditCardService.remove(+id);
  }
}
