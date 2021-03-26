import { ComponentFixture, TestBed } from '@angular/core/testing';

import { MesPedidoComponent } from './mes-pedido.component';

describe('MesPedidoComponent', () => {
  let component: MesPedidoComponent;
  let fixture: ComponentFixture<MesPedidoComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ MesPedidoComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(MesPedidoComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
