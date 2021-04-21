import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CargarEntregasComponent } from './cargar-entregas.component';

describe('CargarEntregasComponent', () => {
  let component: CargarEntregasComponent;
  let fixture: ComponentFixture<CargarEntregasComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CargarEntregasComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(CargarEntregasComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
