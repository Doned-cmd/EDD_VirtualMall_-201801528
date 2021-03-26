import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DiaPedidoComponent } from './dia-pedido.component';

describe('DiaPedidoComponent', () => {
  let component: DiaPedidoComponent;
  let fixture: ComponentFixture<DiaPedidoComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ DiaPedidoComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(DiaPedidoComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
