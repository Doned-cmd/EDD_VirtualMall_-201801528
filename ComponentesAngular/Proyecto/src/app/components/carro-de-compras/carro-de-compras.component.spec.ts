import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CarroDeComprasComponent } from './carro-de-compras.component';

describe('CarroDeComprasComponent', () => {
  let component: CarroDeComprasComponent;
  let fixture: ComponentFixture<CarroDeComprasComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CarroDeComprasComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(CarroDeComprasComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
